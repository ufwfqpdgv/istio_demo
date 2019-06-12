## 项目
这个项目提供了飒漫画的漫画月票数和排行查询的服务

## 项目功能
- 每30分钟加载114服务器上samh.comic_sort_rank数据表中最新的漫画月票信息
- 提供服务,根据漫画id来获得月票数和排行

## 项目结构
- 代码放在src/samh_ticket_rank_0.1下
- recommend文件夹负责离线获取漫画id和月票数及排行的信息和根据漫画id来获取月票信息的函数
- server文件夹负责搭建服务

## 请求url
http://47.99.39.114:8114/ticket/v1?comicId=xxx

## 请求参数说明
comicId, 漫画id

## 返回
- msg,信息码,"ok"则正常
- status, 状态码, "0"则正常
- data, 漫画的月票信息,包括
    - thismonth_ticket_num, string, 这个月的月票数,默认为10-13的整数
    - Rank, string, 这个月的月票排行，默认为1111