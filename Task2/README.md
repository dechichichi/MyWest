
FZU:
爬取福大通知、文件系统
爬取福州大学通知、文件系统

包含发布时间，作者，标题以及正文。
可自动翻页（爬虫可以自动对后续页面进行爬取，而不需要我们指定第几页）
范围：2020年1月1号 - 2021年9月1号（不要爬太多了）。
Bonus:
使用并发爬取，同时给出加速比（加速比：相较于普通爬取，快了多少倍）
搜集每个通知的访问人数
将爬取的数据存入数据库，原生SQL或ORM映射都可以

BiliBili:
爬取Bilibili视频评论
爬取 https://www.bilibili.com/video/BV12341117rG 的全部评论
全部评论，包含子评论
Bonus:
给出Bilibili爬虫检测阈值（请求频率高于这个阈值将会被ban。也可以是你被封时的请求频率）
给出爬取的流程图，使用mermaid或者excalidraw
给出接口返回的json中每个参数所代表的意义'