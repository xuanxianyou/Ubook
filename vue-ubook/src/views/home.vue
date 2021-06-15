<template>
  <el-container>
    <el-header height="100px">图书推荐</el-header>
    <el-main class="book-brief">
      <el-row class="book-name">{{ book.name }}</el-row>
      <el-row :gutter="20">
        <el-col :span="4">
          <div class="grid-content bg-purple">
            <div class="demo-image__preview">
              <el-image class="book-img" :src="url" :preview-src-list="srcList">
              </el-image>
            </div>
          </div>
        </el-col>
        <el-col :span="12">
          <div class="grid-content bg-purple">
            <el-row class="book-info"> 作者：{{book.author}} </el-row>
            <el-row class="book-info"> 出版社：{{book.press}} </el-row>
            <el-row class="book-info"> 出版日期：2016-05 </el-row>
            <el-row class="book-info"> 页数：{{book.page}} </el-row>
            <el-row class="book-info"> 定价：{{book.price}} </el-row>
            <el-row class="book-info"> 丛书： {{book.series}} </el-row>
            <el-row class="book-info"> ISBN: {{book.ISBN}} </el-row>
            <el-row class="book-info">
              豆瓣评分：
              <el-rate
                v-model="book.score"
                disabled
                show-score
                text-color="#ff9900"
                score-template="{value}"
              >
              </el-rate>
            </el-row>
          </div>
        </el-col>
        <el-col :span="8">
          <el-row class="grid-content bg-purple" :gutter="20">
            <el-col :span="1">
              <el-divider direction="vertical"></el-divider>
            </el-col>
            <el-col :span="23">
              <el-row>标签：</el-row>
              <el-tag v-for="tag in book.tags" :key="tag">{{tag}}</el-tag>
            </el-col>
          </el-row>
        </el-col>
      </el-row>
    </el-main>
    <el-main class="book-content">
      <el-row class="title">内容简介</el-row>
      <el-divider></el-divider>
      <p>{{ book.contentBrief }}</p>
    </el-main>
    <el-main class="book-content">
      <el-row class="title">作者简介</el-row>
      <el-divider></el-divider>
      <el-row><p>拉斯·古斯塔夫松Lars Gustafsson（1936-2016）</p></el-row>
      <el-row>
        <p>{{book.authorBrief}}</p>
      </el-row>

    </el-main>
    <el-main class="book-content">
      <el-row class="title">目录</el-row>
      <el-divider></el-divider>
      <!--<p v-for="catalogue in catalogues" v-bind:todo="catalogue"
      v-bind:key="catalogue.id">
          {{book.catalogue}}
      </p>-->
      <div v-html="book.catalogue"></div>
    </el-main>
  </el-container>
</template>

<script>
export default {
  data() {
    return {
      value: 3.7,
      url: "https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg",
      srcList: [
        "https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg",
        "https://fuss10.elemecdn.com/1/8e/aeffeb4de74e2fde4bd74fc7b4486jpeg.jpeg",
      ],
      tagList: ["","success", "info", "warning","danger"],
      book:{
        'name': '',
        'author':'',
        'press':'',
        'publication': '',
        'page': 0,
        'price': 0,
        'series': '',
        'ISBN': '',
        'score': '',
        'tags': [],
        'contentBrief': '',
        'authorBrief': '',
        'catalogue': '',
      }
    };
  },
  created() {
    this.getBook()
  },
  methods:{
    getBook() {
      const { data } = this.$get('/book/getbook',{'userId': 1})
            .then((data) => {
              // 获取图书信息
              if (data.code === 200 && data.success === true) {
                this.book.name = data.book.name
                this.book.author = data.book.author
                this.book.press = data.book.press
                this.book.publication = data.book.publication
                this.book.page = data.book.page
                this.book.price = data.book.price
                this.book.series = data.book.series
                this.book.ISBN = data.book.isbn
                this.book.score = parseFloat(data.book.score)/2
                this.book.tags = data.book.tags
                this.book.contentBrief = data.book.content_brief
                this.book.authorBrief = data.book.author_brief
                this.book.catalogue = this.transferToHTML(data.book.catalogue)
                this.$message({
                  type: 'success',
                  message: data.message
                })
              } else {
                this.$message({
                  type: 'error',
                  message: data.message
                })
                return false
              }
            }).catch(() => {
              
            })
          console.log(data)
    },
    transferToHTML(string){
      //替换所有的换行符
      string = string.replace(/\r\n/g,"<br>")
      string = string.replace(/\n/g,"<br>");
      
      //替换所有的空格（中文空格、英文空格都会被替换）
      string = string.replace(/\s/g,"&nbsp;");
      
      //输出转换后的字符串
      console.log(string);
      return string
    }
  }
};
</script>

<style lang="scss" scoped>
body {
  padding: 10px 30px;
}
.el-header {
  background-color: #b3c0d1;
  color: #333;
  text-align: center;
  line-height: 100px;
  font-size: 20px;
  border-radius: 5px;
  margin: 5px 20px;
}
.book-brief {
  background-color: #e9eef3;
  color: #000;
  border-radius: 5px;
  margin: 5px 20px;
  .book-name {
    font-size: 25px;
    margin: 10px 0;
  }
}
.book-content{
  background-color: #e9eef3;
  color: #000;
  border-radius: 5px;
  margin: 5px 20px;
  .title{
    font-weight: bold;
  }
}
.book-img {
  width: 200px;
  height: 300px;
}
.book-info {
  line-height: 35px;
}
.el-divider--vertical {
  display: inline-block;
  width: 1px;
  height: 300px;
  margin: 0 8px;
  vertical-align: middle;
  position: relative;
}
.el-divider{
  margin: 10px 0;
}
.el-tag {
  margin: 10px;
}
p{
  text-indent: 2em;
}
</style>