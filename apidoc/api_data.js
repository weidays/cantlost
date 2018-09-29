define({ "api": [  {    "type": "",    "url": "POST",    "title": "http://aip.fixme.vip/v0/lostinfo",    "group": "lostinfo",    "parameter": {      "fields": {        "Parameter": [          {            "group": "Parameter",            "optional": false,            "field": "page",            "description": "<p>int 第几页</p>"          },          {            "group": "Parameter",            "optional": false,            "field": "page_size",            "description": "<p>int 每页条数</p>"          },          {            "group": "Parameter",            "optional": false,            "field": "category_id",            "description": "<p>int 分类id</p>"          }        ]      }    },    "examples": [      {        "title": "请求示例",        "content": "{\n\t\t  \"page\": 1,\n\t\t  \"page_size\": 15,\n\t\t  \"category_id\": 18\n\t\t}",        "type": "json"      },      {        "title": "返回数据示例",        "content": "{\n\t    \"count\": 1,\n\t    \"current:\": 1,\n\t    \"data\": [\n\t        {\n\t            \"id\": 0,\n\t            \"publish_at\": null,\n\t            \"updated_at\": null,\n\t            \"user_id\": 0,\n\t            \"member_info\": {\n\t                \"id\": 1,\n\t                \"registed_at\": \"2018-09-30T08:09:57.018+07:00\",\n\t                \"updated_at\": \"2018-09-30T08:10:01.942+07:00\",\n\t                \"pic_url\": \"https://p1.4499.cn/touxiang/UploadPic/2013-4/27/201304271944118052.jpg\",\n\t                \"nick_name\": \"老妖怪\",\n\t                \"real_name\": \"l老妖\",\n\t                \"phone_number\": \"13808998988\",\n\t                \"gender\": \"2\",\n\t                \"birthDay\": \"2018-09-30T08:14:34.085+07:00\",\n\t                \"login_name\": \"weidays\",\n\t                \"password\": \"goodboy\"\n\t            },\n\t            \"category_info\": null,\n\t            \"category_id\": 0,\n\t            \"title\": \"\",\n\t            \"content\": \"\",\n\t            \"like_num\": 0,\n\t            \"dislike_num\": 0,\n\t            \"comment_num\": 0,\n\t            \"share_num\": 0\n\t        }\n\t    ],\n\t    \"message\": \"SUCCESS\",\n\t    \"page_size\": 15,\n\t    \"status\": 200\n\t}",        "type": "json"      }    ],    "version": "0.0.0",    "filename": "routers/routers.go",    "groupTitle": "lostinfo",    "name": "Post"  }] });
