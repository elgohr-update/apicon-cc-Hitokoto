package main

import (
	"io/ioutil"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	jsoniter "github.com/json-iterator/go"
	"github.com/wuhan005/gadget"
	log "unknwon.dev/clog/v2"
)

func init() {
	_ = log.NewConsole(100)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	categoriesRaw, err := ioutil.ReadFile("categories.json")
	if err != nil {
		log.Fatal("Failed to read categories.json: %v", err)
	}
	var categories []Category
	err = jsoniter.Unmarshal(categoriesRaw, &categories)
	if err != nil {
		log.Fatal("Failed to unmarshal categories.json: %v", err)
	}
	log.Trace("Parse categories.json, %d categories loaded.", len(categories))

	sentencesDict := map[string][]Sentence{}
	for _, category := range categories {
		sentencesRaw, err := ioutil.ReadFile(category.Path)
		if err != nil {
			log.Fatal("Failed to read %s: %v", category.Path, err)
		}
		var sentences []Sentence
		err = jsoniter.Unmarshal(sentencesRaw, &sentences)
		if err != nil {
			log.Fatal("Failed to unmarshal %s: %v", category.Path, err)
		}

		// 装填数据
		if sentencesDict[category.Key] == nil {
			sentencesDict[category.Key] = make([]Sentence, 0, len(sentences))
		}
		for _, sen := range sentences {
			sentencesDict[category.Key] = append(sentencesDict[category.Key], sen)
		}

		log.Trace("Read %s, %d sentences loaded.", category.Path, len(sentences))
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		// 句子类型
		sentenceTypes, ok := c.GetQueryArray("c")
		if !ok {
			sentenceTypes = []string{"a", "b", "c", "e", "f", "g", "h", "i", "j", "k", "l"}
		}
		for i, t := range sentenceTypes {
			// 不存在的类型，当做动画 (a) 处理
			if _, ok := sentencesDict[t]; !ok {
				sentenceTypes[i] = "a"
			}
		}

		encode := c.DefaultQuery("encode", "text") // 编码

		// 选择句子类型
		sType := sentenceTypes[rand.Intn(len(sentenceTypes))]
		// 选择句子
		dict := sentencesDict[sType]
		sentence := dict[rand.Intn(len(dict))]

		switch encode {
		case "text":
			c.Render(200, render.Data{
				ContentType: "text/plain; charset=utf-8",
				Data:        []byte(sentence.Hitokoto),
			})
		case "js":
			callback := c.DefaultQuery("callback", "hitokoto")
			if callback == "" {
				c.Render(200, render.JSON{Data: sentence})
				return
			}
			c.Render(200, render.JsonpJSON{Callback: callback, Data: sentence})
		default:
			// 默认为 JSON
			c.JSON(200, sentence)
		}
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(gadget.MakeErrJSON(404, 40400, "not found"))
	})
	log.Fatal("Failed to run web server: %v", r.Run(":8080"))
}
