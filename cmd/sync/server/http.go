package server

import (
	"github.com/TurboHsu/munager/cmd/sync/structure"
	"github.com/TurboHsu/munager/util/logging"
	"github.com/gin-gonic/gin"
)

func ListenAndServe(addr string) {
	r := gin.Default()
	r.POST("/api/handshake", handshakeAPIHandler)
	r.POST("/api/get-list", listAPIHandler)
	r.POST("/api/get-file", fileAPIHandler)

	logging.HandleErr(r.Run(addr))
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
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "bad request",
		})
		return
	}

	var diff []structure.FileInfo

	// Find all files in local
	local := getFiles(ServerCommand.Flag("path").Value.String())

	// Diff all the files
	for _, l := range local {
		flag := false
		for _, r := range result.Files {
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
	doesTerminateBroadcast, err := ServerCommand.Flags().GetBool("keep-broadcasting")
	logging.HandleErr(err)
	if doesTerminateBroadcast {
		TerminateBroadcast = true
	}

	registerUser(result.Fingerprint)

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
	})
}
