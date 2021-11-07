package music

import "github.com/gin-gonic/gin"

type MusicInfo struct {
	id     int64
	Name   string
	Singer string
	Cover  string
	Time   string
	Url    string
}

func MusicList(c *gin.Context) {
	musics := []MusicInfo{
		{Name: "反方向的钟", Singer: "周杰伦", Cover: "", Time: "4:30", Url: ""},
		{Name: "幻听", Singer: "许嵩", Cover: "", Time: "5:20", Url: ""},
	}

	c.JSON(200, musics)
}

func AddMusic(c *gin.Context) {
	c.JSON(200, gin.H{"todo": "todo"})
}

func PlayMusic(c *gin.Context) {
	id := c.Param("id")

	c.String(200, "Playing music %s", id)
}

func MusicHome(c *gin.Context) {
	c.HTML(200, "music_index.html", "")
}
