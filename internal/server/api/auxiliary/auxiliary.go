package auxiliary

import (
	"net/http"
	"time"

	"github.com/cyberninja89/website/internal/consts"

	"github.com/gin-gonic/gin"
	"github.com/go-nacelle/process"
)

// Linkable variables for arguments at build
var (
	Version        = "No Version Provided"
	BuildTimestamp = "No Build Time"
	GitHash        = "No Git Information"
)

type versionResponse struct {
	Version   string `json:"Version"`
	BuildTime string `json:"Build Timestamp"`
	GitHash   string `json:"Git Hash"`
}

func (r *Router) Version(c *gin.Context) {
	resp := versionResponse{
		Version:   Version,
		BuildTime: BuildTimestamp,
		GitHash:   GitHash,
	}

	c.IndentedJSON(http.StatusOK, resp)
}

type statusResponse struct {
	AppName string           `json:"App Name"`
	Time    string           `json:"Timestamp"`
	Status  string           `json:"Status"`
	Reasons []process.Reason `json:"Reasons,omitempty"`
}

func (r *Router) Status(c *gin.Context) {
	resp := statusResponse{
		AppName: consts.ProcessName,
		Time:    time.Now().Format(consts.TimeFormat),
		Status:  "NOT READY",
		Reasons: r.Health.Reasons(),
	}
	if len(r.Health.Reasons()) == 0 {
		resp.Status = "READY"
	}

	c.IndentedJSON(http.StatusOK, resp)
}
