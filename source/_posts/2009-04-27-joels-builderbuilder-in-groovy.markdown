---
date: '2009-04-27 21:54:07'
layout: post
slug: joels-builderbuilder-in-groovy
status: publish
title: Joel's BuilderBuilder in Groovy
wordpress_id: '149'
categories:
- functional-programming
- groovy
---

[Joel Neely](http://joelneely.wordpress.com/) started [a series](http://joelneely.wordpress.com/2009/04/25/builderbuilder-the-agenda/) [of posts](http://joelneely.wordpress.com/2009/04/27/builderbuilder-the-task/) over the weekend detailing a proposed exploration of what Functional Programming means "to me as a practicing OO software developer?" The task at hand is to look at the generation of Data Transfer Objects which include a static inner class that functions as a builder. Since I'm exploring both Clojure and Scala right now, Joel has asked me to play along.

After reading the task definition today, I couldn't resist taking a crack at the problem in Groovy, especially since tackling a Builder DSL is not something I've gone after before.

So, here's how I'll use my DTOBuilder to produce the output Joel described:


    
``` java   
def bldr = new DtoBuilder()

println bldr.build {
	packageName 'edu.bogusu.registration'
	name 'Student'
	field(name:'id', type:'String')
	field(name:'firstName', type:'String')
	field(name:'lastName', type:'String')
	field(name:'hoursEarned', type:'int')
	field(name:'gpa', type:'float')
}
```    



Pretty concise, eh? [Check out the implementation of it at GitHub](http://github.com/mstine/BuilderBuilder/blob/1008235f88177910eb94af165ade3aedae2955a3/src/DtoBuilder.groovy). I'll be posting all of my code related to this series at this location (and hopefully Joel will join in as well).

One thing you'll notice is that I'm still operating primarily in OO style. My next task is to refactor this code, still written in Groovy, but using as much functional-style as I can squeeze out of the language. Until next time...
