---
date: '2009-04-25 23:09:11'
layout: post
slug: groovy-post-to-a-url
status: publish
title: 'Groovy: Post to a URL'
wordpress_id: '146'
categories:
- CodeProject
- groovy
---

If you've ever worked with the brilliant [Recaptcha](http://recaptcha.net/) service, you'll know that their REST API requires an HTTP POST rather than a GET. As I had only used GET requests thus far, I googled around and found a pretty easy solution:


    
    
    private def evaluateCaptcha(def remoteIp, def challenge, def response) {
        def config = recaptchaService.getRecaptchaConfig()
    
        def urlString = "http://api-verify.recaptcha.net/verify"
        def queryString = "privatekey=${config.recaptcha.privateKey}&remoteip=${remoteIp}&challenge=${challenge}&response=${URLEncoder.encode(response)}"
    
        def url = new URL(urlString)
        def connection = url.openConnection()
        connection.setRequestMethod("POST")
        connection.doOutput = true
    
        def writer = new OutputStreamWriter(connection.outputStream)
        writer.write(queryString)
        writer.flush()
        writer.close()
        connection.connect()
    
        def recaptchaResponse = connection.content.text
        log.debug(recaptchaResponse)
    
        recaptchaResponse.startsWith("true")
    }
    



I have to credit [Justin Spradlin](http://www.fiascode.com/programming/putting-google-finance-to-rest-with-groovy/) for the code that ultimately got me here. Consider this a +1.
