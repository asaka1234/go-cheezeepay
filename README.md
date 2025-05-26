文档
=============
https://pay-apidoc-en.cheezeebit.com/#cheezee-pay-api

鉴权
==============
1. rsa privateKey私钥加密算签名, rsa publicKey公钥解密验证签名
2. 对请求参数算了一个sign签名, 随后作为json里一个字段一起发

回调地址
==============


Comment
===============
1. only support deposit
2. 所有接口都是 application/json 格式的
