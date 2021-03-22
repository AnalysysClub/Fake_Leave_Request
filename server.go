package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("./templates/*.html")
	r.GET("/form", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "form.html", gin.H{})

	})
	r.GET("/leave", func(ctx *gin.Context) {
		//这些都是用户输入的部分
		name := ctx.Query("name")
		student_id := ctx.Query("student_id")
		reason := ctx.Query("reason")
		coach_name := ctx.Query("coach_name")
		tele_phone := ctx.Query("tele_phone")
		identity_card := ctx.Query("identity_card")
		emergency_contant := ctx.Query("emergency_contant")
		emergency_tele := ctx.Query("emergency_tele")
		//用户输入部分截止
		//生成现在的时间
		time_now := time.Now().Format("2006-01-02 15:04")
		time_post := []byte(time_now)
		time_post[9] += 1
		time_postpone := string(time_post)
		//时间生成完毕
		//审批编号随机生成
		approval_num :=
		//审批编号随机生成结束
		location := ctx.Query("location")
		ctx.HTML(http.StatusOK, "detail.html", gin.H{
			"name":              name,
			"student_id":        student_id,
			"reason":            reason,
			"coach_name":        coach_name,
			"identity_card":     identity_card,
			"tele_phone":        tele_phone,
			"emergency_tele":    emergency_tele,
			"emergency_contant": emergency_contant,
			"location":          location,
			"time_now":          time_now,
			"time_postpone":     time_postpone,
		})
	})
	r.Run(":8081")
}
