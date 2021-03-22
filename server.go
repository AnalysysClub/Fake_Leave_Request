package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
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
		identity_card := ctx.Query("identity_card") //身份证ID
		emergency_contant := ctx.Query("emergency_contant")
		emergency_tele := ctx.Query("emergency_tele")
		//用户输入部分截止
		//随机生成身份证:
		//identity_card := rand_string(18) //还是取消了把，感觉生成的太离谱了
		//生成现在的时间
		time_now_real := time.Now().Format("2006-01-02 15:04")
		time_now_conv := []byte(time_now_real)
		time_now_conv[12] -= 1
		time_now := string(time_now_conv)
		time_post := []byte(time_now_real)
		time_post[9] += 1
		time_postpone := string(time_post)
		//时间生成完毕
		//审批编号随机生成

		approval_num := time_now[:4] + time_now[5:7] + time_now[8:10] + rand_string(12)
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
			"approval_num":      approval_num,
		})
	})
	r.Run("192.168.3.102:8081")
}

func rand_string(length int) string {
	result := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		num := rand.Intn(10)
		result = result + strconv.Itoa(num)
	}
	return result
}
