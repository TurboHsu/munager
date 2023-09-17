package server

import (
	"os"
	"time"

	"github.com/TurboHsu/munager/cmd/sync/structure"
	"github.com/TurboHsu/munager/cmd/sync/utils"
	fileprocessing "github.com/TurboHsu/munager/util/file"
	"github.com/TurboHsu/munager/util/logging"
	"github.com/gin-gonic/gin"
)

func ListenAndServe(addr string) {
	// Set release mode
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.POST("/api/handshake", handshakeAPIHandler)
	r.POST("/api/get-list", listAPIHandler)
	r.POST("/api/get-file", fileAPIHandler)
	r.POST("/api/get-checksum", checksumAPIHandler)
	r.POST("/api/suicide", suicideHandler)

	logging.HandleErr(r.Run(addr))
}

func suicideHandler(c *gin.Context) {
	var result structure.Suicide
	err := c.ShouldBindJSON(&result)
	if err != nil {
		logging.HandleErr(err)
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "bad request",
		})
		return
	}

	// Destorys the user
	killUser(result.Fingerprint)

	// Log it
	logging.Info("User " + result.Fingerprint + " destructed.")

	// Checks whether this server need to be running
	keepBrocasting, err := ServerCommand.Flags().GetBool("keep-broadcasting")
	if err != nil {
		logging.HandleErr(err)
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "bad request",
		})
		return
	}

	// Checks whether needs to take a break
	go func() {
		time.Sleep(1 * time.Second)
		if !keepBrocasting && len(UserBase) == 0 {
			// Time for a break
			os.Exit(0)
		}
	}()

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
	})
}

func checksumAPIHandler(c *gin.Context) {
	var result structure.FileServeRequest
	err := c.ShouldBindJSON(&result)
	if err != nil {
		logging.HandleErr(err)
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "bad request",
		})
		return
	}

	user := getUser(result.Fingerprint)
	if user == nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "bad request",
		})
		return
	}

	// Check if the file is in user's access list
	flag := false
	var file structure.FileInfo
	for _, f := range user.AccessList {
		if f.PathBase == result.PathBase && f.Extension == result.Extension {
			flag = true
			file = f
			break
		}
	}

	if !flag {
		c.JSON(400, gin.H{
			"code": 403,
			"msg":  "forbidden",
		})
		return
	}

	// Serves the file
	servingPath := ServerCommand.Flag("path").Value.String() + file.PathBase + "." + file.Extension
	f, err := os.OpenFile(servingPath, os.O_RDONLY, 0644)
	if err != nil {
		logging.HandleErr(err)
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  "internal server error",
		})
		return
	}
	defer f.Close()
	checksum := fileprocessing.CalculateSHA1(f)
	c.JSON(200, structure.ChecksumStruct{
		Checksum: checksum,
	})
}

// This function serves file to client
func fileAPIHandler(c *gin.Context) {
	var result structure.FileServeRequest
	err := c.ShouldBindJSON(&result)
	if err != nil {
		logging.HandleErr(err)
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "bad request",
		})
		return
	}

	user := getUser(result.Fingerprint)
	if user == nil {
		c.JSON(403, gin.H{
			"code": 403,
			"msg":  "forbidden",
		})
		return
	}

	// Check if the file is in user's access list
	flag := false
	var file structure.FileInfo
	for _, f := range user.AccessList {
		if f.PathBase == result.PathBase && f.Extension == result.Extension {
			flag = true
			file = f
			break
		}
	}

	if !flag {
		c.JSON(403, gin.H{
			"code": 403,
			"msg":  "forbidden",
		})
		return
	}

	// Serves the file
	servingPath := ServerCommand.Flag("path").Value.String() + file.PathBase + "." + file.Extension
	logging.Info("Serving file " + servingPath + " to " + c.ClientIP() + "...")
	c.File(servingPath)

	// Deletes the file from user's access list
	user.DeleteFile(result.PathBase)
}

// This function handles file list from client, and give out the difference between local and remote
func listAPIHandler(c *gin.Context) {
	var result structure.ListRequest
	err := c.ShouldBindJSON(&result)
	if err != nil {
		logging.HandleErr(err)
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "bad request",
		})
		return
	}

	user := getUser(result.Fingerprint)
	if user == nil {
		c.JSON(403, gin.H{
			"code": 403,
			"msg":  "forbidden",
		})
		return
	}

	var diff []structure.FileInfo

	// Find all files in local
	local := utils.GetFiles(ServerCommand.Flag("path").Value.String())
	local = utils.FilterValidFiles(local)

	// Diff all the files
	for _, l := range local {
		flag := false
		for _, r := range result.Files {
			// Checks PathBase only, because client and server may have different file format
			if l.PathBase == r.PathBase {
				flag = true
				break
			}
		}
		if !flag {
			diff = append(diff, l)
		}
	}

	// Grant all the file access to the user
	user.AccessList = diff

	reply := structure.ListResponse{
		Files: diff,
	}

	c.JSON(200, reply)
}

// This function handles handshake with client, it creates a user instance with its given fingerprint
func handshakeAPIHandler(c *gin.Context) {
	var result structure.Handshake
	err := c.ShouldBindJSON(&result)
	if err != nil {
		logging.HandleErr(err)
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "bad request",
		})
		return
	}

	if result.MagicWord != structure.HandshakeMagicWord {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "bad request",
		})
		return
	}

	// Terminate broadcasting
	doesNotTerminateBroadcast, err := ServerCommand.Flags().GetBool("keep-broadcasting")
	logging.HandleErr(err)
	if !doesNotTerminateBroadcast {
		logging.Info("Client " + c.ClientIP() + " handshaked with server, terminating broadcasting...")
		TerminateBroadcast = true
	}

	registerUser(result.Fingerprint)

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
	})
}
