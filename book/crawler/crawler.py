import asyncio
import time

import utils
from time import sleep

from urllib import parse
from pyppeteer import launch

BaseURL = "https://book.douban.com/tag/"
tags = ["数学", "编程", "算法", "科技", "web", "神经网络"]
Headers = {

}


def screen_size():
    # 使用tkinter获取屏幕大小
    import tkinter
    tk = tkinter.Tk()
    width = tk.winfo_screenwidth()
    height = tk.winfo_screenheight()
    tk.quit()
    return width, height


async def generate_tag_url():
    """
    根据tag生成url
    :return:
    """
    tag_url_list = []
    for tag in tags:
        escape = parse.quote(tag)
        tag_url_list.append(BaseURL + escape)
    return tag_url_list


class Crawler:
    def __init__(self):
        self.browser = None
        self.page = None

    async def GenerateURL(self, tag_url_list):
        url_list = []
        if self.page is None:
            self.page = await self.browser.newPage()
        for tag_url in tag_url_list:
            # 遍历所有tag
            for i in range(0, 5000, 20):
                # 遍历标签中的所有页
                pageURL = "{}?start={}&type=T".format(tag_url, i)
                await self.page.goto(pageURL)
                book_url_list = await self.page.xpath("//*[@id='subject_list']/ul/li/div[2]/h2/a")
                score_list = await self.page.xpath("//*[@id='subject_list']/ul/li/div[2]/div[2]")
                if len(book_url_list) == 0:
                    break
                for book_url, score in zip(book_url_list, score_list):
                    url = await (await book_url.getProperty('href')).jsonValue()
                    s = await (await score.getProperty('textContent')).jsonValue()
                    num = await utils.GetEvaluateNum(s)
                    score = await utils.GetScore(s)
                    if num > 100 and score > 7.0:
                        url_list.append(url)
                        print(url)
                        await self.RequestURL(url)
                num = await self.browser.pages()
                if len(num) >= 10:
                    await asyncio.sleep(5)
                sleep(4)
        return url_list

    async def RequestURL(self, url):
        page = await self.browser.newPage()
        await page.goto(url)
        await self.ParsePage(page)
        await page.close()

    async def ParsePage(self, page):
        # parse page
        name_l = await page.xpath("//div[@id='wrapper']/h1/span")
        cover_l = await page.xpath("//div[@id='mainpic']/a/img")
        info_l = await page.xpath("//div[@id='info']")
        score_l = await page.xpath("//div[@id='interest_sectl']/div/div[2]/strong")
        content_brief_l = await page.xpath("//div[@id='link-report']/div[1]/div/p")
        author_brief_l = await page.xpath("//div[@id='content']/div/div[1]/div[3]/div[3]/div/div")
        tags_l = await page.xpath("//div[@id='db-tags-section']/div")
        for name_e, cover_e, info_e, score_e, content_brief_e, author_brief_e, tags_e in zip(name_l, cover_l, info_l,
                                                                                             score_l, content_brief_l,
                                                                                             author_brief_l, tags_l):
            name = await (await name_e.getProperty('textContent')).jsonValue()
            cover = await (await cover_e.getProperty('src')).jsonValue()
            info = await (await info_e.getProperty('textContent')).jsonValue()
            score = await (await score_e.getProperty('textContent')).jsonValue()
            content_brief = await (await content_brief_e.getProperty('textContent')).jsonValue()
            author_brief = await (await author_brief_e.getProperty('textContent')).jsonValue()
            tags = await (await tags_e.getProperty('textContent')).jsonValue()

            print(name)
            print(cover)
            print(info)
            print(score)
            print(content_brief)
            print(author_brief)
            print(tags)

    def Save(self):
        pass

    async def run(self):
        self.browser = await launch(
            {'headless': False, "userDataDir": "./userDataDir", 'dumpio': True, 'autoClose': False,
             'args': ['--no-sandbox']})
        page = await self.browser.newPage()
        width, height = screen_size()
        await page.setViewport({
            "width": width,
            "height": height
        })
        tag_url_list = await generate_tag_url()
        await self.GenerateURL(tag_url_list)
        # await self.RequestURL("https://book.douban.com/subject/35426737/")
        await self.browser.close()


if __name__ == '__main__':
    crawler = Crawler()
    # crawler.GenerateURL()
    asyncio.get_event_loop().run_until_complete(crawler.run())
