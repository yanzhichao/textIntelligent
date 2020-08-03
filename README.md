## 文本智能处理服务

### 1、语种识别

调用方式如下
```
curl -H "Content-Type:application/json" -H "Data_Type:msg" -X POST --data '{"context": "test hello world"}' http://127.0.0.1:8080//v1/intelligent/text/getLanguageDetection
```