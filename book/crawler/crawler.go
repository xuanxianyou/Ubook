package crawler

import (
	"GoProject/WebProject/MicroService/UserBook/book/model"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/antchfx/htmlquery"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var(
	BaseURL="https://book.douban.com/tag/"
	tags=[]string{"数学","编程","算法","科技","web","神经网络"}
	URLChannel=make(chan string, 20)
	BodyChannel=make(chan string,20)
	BookChannel=make(chan model.Book,20)
	Header=map[string]string{
		"Host":"book.douban.com",
		"User-Agent":"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.77 Safari/537.36",
		"Cookie":"bid=djfnYOWMNUc; _vwo_uuid_v2=D4DCD9B52EFDE00F84FCBC333CA0917F1|355bebb2d5e5c84ef77543ac845decc0; gr_user_id=d8a97a05-6340-4137-bc3d-4f8b7ce1bbfa; __yadk_uid=H397XuChJ39LJoCd5ApqiJQ2meW83GJS; __gads=ID=a5b649fb999c7bfe-2213d13e1fc90067:T=1622648045:RT=1622648045:S=ALNI_MbPWXS6nmOHXoI_WDqT6NU_ynsHIQ; douban-fav-remind=1; ct=y; viewed=\"19952400_30136932_26829016_35196328_26759508_35006892_26941639_3676140_33444476_30293801\"; dbcl2=\"199515848:P3gu4jB46kk\"; push_noty_num=0; push_doumail_num=0; ap_v=0,6.0; ck=FyVv; __utmz=30149280.1623236333.18.14.utmcsr=baidu|utmccn=(organic)|utmcmd=organic; __utmc=30149280; __utma=30149280.134614742.1622648038.1623230538.1623236333.18; __utmz=81379588.1623236333.18.14.utmcsr=baidu|utmccn=(organic)|utmcmd=organic; __utma=81379588.918737593.1622648038.1623230538.1623236333.18; __utmc=81379588; _pk_ref.100001.3ac3=%5B%22%22%2C%22%22%2C1623236333%2C%22https%3A%2F%2Fwww.baidu.com%2Flink%3Furl%3D1hcJi01O28Lk59lj1jSOQcKTj5COunxfj2iRHoaxHCbW5fZGjsns4l0Z5FqSXH0Z%26wd%3D%26eqid%3Db5e4f547000016d00000000360c09eea%22%5D; _pk_ses.100001.3ac3=*; __utmt_douban=1; __utmt=1; gr_session_id_22c937bbd8ebd703f2d8e9445f7dfd03=bd91b816-d738-40ee-92e7-c999ffe623f7; gr_cs1_bd91b816-d738-40ee-92e7-c999ffe623f7=user_id%3A1; gr_session_id_22c937bbd8ebd703f2d8e9445f7dfd03_bd91b816-d738-40ee-92e7-c999ffe623f7=true; _pk_id.100001.3ac3=ec075db90832d11a.1622424178.19.1623238947.1623234375.; __utmb=30149280.12.10.1623236333; __utmb=81379588.12.10.1623236333",
	}
)

type Crawler struct {
	Client http.Client
}

// 创建爬虫
func NewCrawler()*Crawler{
	client:=http.Client{}

	return &Crawler{
		Client: client,
	}
}
// 创建URL生成器
func (c *Crawler)GenerateURL(){
	for index,tag:=range tags{
		escape:=url.QueryEscape(tag)
		tagURL:=BaseURL+escape
		fmt.Println(index,tagURL)
		for i:=0;;i+=20{
			pageURL:=fmt.Sprintf("%s?start=%d&type=T",tagURL,i)
			req,err:=http.NewRequest(http.MethodGet,pageURL,nil)
			if err!=nil{
				log.Fatal(err)
			}
			// 设置请求头
			req.Header.Set("Host",Header["Host"])
			req.Header.Set("Cookie",Header["Cookie"])
			req.Header.Set("User-Agent",Header["User-Agent"])
			resp,err:=c.Client.Do(req)
			if err!=nil{
				log.Fatal(err)
			}
			body ,err :=ioutil.ReadAll(resp.Body)
			if err!=nil{
				log.Fatal(err)
			}
			if resp.StatusCode==404{
				break
			}
			// fmt.Println(string(body))
			// 解析
			// 如果列表为空，break，否则，添加到URLChannel中
			// xpath=`//*[@id="subject_list"]//h2/a`
			dom,err:=goquery.NewDocumentFromReader(strings.NewReader(string(body)))
			if err!=nil{
				log.Println(err.Error())
			}
			var hrefs []string
			if dom!=nil{
				dom.Find("div>h2>a").Each(func(i int, selection *goquery.Selection) {
					href,_:=selection.Attr("href")
					hrefs = append(hrefs, href)
				})
			}
			if len(hrefs)==0{
				break
			}
			for _,href:=range hrefs{
				fmt.Println(href)
				URLChannel<-href
			}
				fmt.Println(pageURL)
		}
		time.Sleep(5*time.Second)
	}
}
// 请求URL
func (c *Crawler)RequestURL(url string){
	client:=c.Client
	req,err:=http.NewRequest(http.MethodGet,url,nil)
	if err!=nil{
		log.Fatal(err)
	}
	req.Header.Set("Host",Header["Host"])
	req.Header.Set("Cookie",Header["Cookie"])
	req.Header.Set("User-Agent",Header["User-Agent"])
	resp,err:=client.Do(req)
	if err!=nil{
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body,err:=ioutil.ReadAll(resp.Body)
	if err!=nil{
		log.Fatal(err)
	}
	dom,err:=goquery.NewDocumentFromReader(strings.NewReader(string(body)))
	if err!=nil{
		log.Println(err.Error())
	}
	var book model.Book
	if dom!=nil{
		dom.Find("div>h1>span").Each(func(i int, selection *goquery.Selection) {
			fmt.Println("书名：")
			book.Name = selection.Text()
			fmt.Println(book.Name)
		})
		dom.Find("#mainpic > a > img").Each(func(i int, selection *goquery.Selection) {
			fmt.Println("封面：")
			fmt.Println(book.Cover)
			book.Cover,_ = selection.Attr("src")
		})
		dom.Find("div>span:nth-child(1)>a").Each(func(i int, selection *goquery.Selection) {
			fmt.Println("作者：")
			book.Author += selection.Text()+" "
			fmt.Println(book.Author)
		})
		dom.Find("#info").Each(func(i int, selection *goquery.Selection) {
			bookInfo:= selection.Text()
			fmt.Println(bookInfo)
		})
		dom.Find("#interest_sectl > div > div.rating_self.clearfix > strong").Each(func(i int, selection *goquery.Selection) {
			score:=selection.Text()
			fmt.Println(score)
		})
		dom.Find("#link-report > div:nth-child(1) > div > p").Each(func(i int, selection *goquery.Selection) {
			content_brief:=selection.Text()
			fmt.Println(content_brief)
		})

	}
}
// 解析Response
func Parse(url string){
	client:=http.Client{}
	req,err:=http.NewRequest(http.MethodGet,url,nil)
	if err!=nil{
		log.Fatal(err)
	}
	req.Header.Set("Host",Header["Host"])
	req.Header.Set("Cookie",Header["Cookie"])
	req.Header.Set("User-Agent",Header["User-Agent"])
	resp,err:=client.Do(req)

	if resp!=nil{
		defer resp.Body.Close()
		doc,_ := htmlquery.Parse(resp.Body)
		name := htmlquery.FindOne(doc, "//div[@id='wrapper']/h1/span")
		cover :=htmlquery.FindOne(doc,"//div[@id='mainpic']/a/img")
		info := htmlquery.FindOne(doc, "//div[@id='info']")
		score := htmlquery.FindOne(doc, "//div[@id='interest_sectl']/div/div[2]/strong")
		contentBrief := htmlquery.FindOne(doc, "//div[@id='link-report']/div[1]/div/p")
		authorBrief := htmlquery.FindOne(doc, "//div[@id='content']/div/div[1]/div[3]/div[3]/div/div")
		tags := htmlquery.FindOne(doc, "//div[@id='db-tags-section']/div")
		fmt.Printf("%s\n",htmlquery.InnerText(name))
		fmt.Println(htmlquery.SelectAttr(cover,"src"))
		fmt.Printf("%s\n",htmlquery.InnerText(info))
		fmt.Printf("%s\n",htmlquery.InnerText(score))
		fmt.Printf("%s\n",htmlquery.InnerText(contentBrief))
		fmt.Printf("%s\n",htmlquery.InnerText(authorBrief))
		fmt.Printf("%s\n",htmlquery.InnerText(tags))
	}


}
// 存储Book

func (c *Crawler)Request(url string, method string){
	client:=c.Client
	req,err:=http.NewRequest(method,url,nil)
	if err!=nil{
		log.Fatal(err)
	}
	req.Header.Set("Host","book.douban.com")
	req.Header.Set("Cookie","bid=djfnYOWMNUc; _vwo_uuid_v2=D4DCD9B52EFDE00F84FCBC333CA0917F1|355bebb2d5e5c84ef77543ac845decc0; gr_user_id=d8a97a05-6340-4137-bc3d-4f8b7ce1bbfa; __yadk_uid=H397XuChJ39LJoCd5ApqiJQ2meW83GJS; __gads=ID=a5b649fb999c7bfe-2213d13e1fc90067:T=1622648045:RT=1622648045:S=ALNI_MbPWXS6nmOHXoI_WDqT6NU_ynsHIQ; douban-fav-remind=1; viewed=\"34926034_33476106_35273146_35362277_35451769_35292523_35450833_34933712\"; ap_v=0,6.0; _pk_ref.100001.3ac3=%5B%22%22%2C%22%22%2C1622981680%2C%22https%3A%2F%2Fwww.baidu.com%2Flink%3Furl%3D-GlIKm4_Sh4T-xph1fnpHfMRXcQsKY0TO2vunU6s8rLaKwchmmgoKKsT9c_-_Sz8%26wd%3D%26eqid%3Dffe9f98e00137e660000000360bcbc2d%22%5D; _pk_ses.100001.3ac3=*; __utma=30149280.134614742.1622648038.1622786862.1622981681.6; __utmc=30149280; __utmz=30149280.1622981681.6.6.utmcsr=baidu|utmccn=(organic)|utmcmd=organic; __utmt_douban=1; __utma=81379588.918737593.1622648038.1622786862.1622981681.6; __utmc=81379588; __utmz=81379588.1622981681.6.6.utmcsr=baidu|utmccn=(organic)|utmcmd=organic; __utmt=1; gr_session_id_22c937bbd8ebd703f2d8e9445f7dfd03=92e485f7-4067-4874-a57a-e2cc4ec9fb55; gr_cs1_92e485f7-4067-4874-a57a-e2cc4ec9fb55=user_id%3A0; gr_session_id_22c937bbd8ebd703f2d8e9445f7dfd03_92e485f7-4067-4874-a57a-e2cc4ec9fb55=true; _pk_id.100001.3ac3=ec075db90832d11a.1622424178.7.1622981881.1622787023.; __utmb=30149280.10.10.1622981681; __utmb=81379588.10.10.1622981681")
	req.Header.Set("User-Agent","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.77 Safari/537.36")
	resp,err:=client.Do(req)
	if err!=nil{
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body,err:=ioutil.ReadAll(resp.Body)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(string(body))
}

func (c *Crawler)Parse(body string)model.Book{
	return model.Book{}
}


func Generator(c chan int,num int) {
	defer close(c) //生成结束后关闭channel
	for i := 0; i < num; i++ {
		c <- i
	}
}
