---
date: '2009-06-05 02:00:59'
layout: post
slug: jug-usa-meeting-with-james-gosling-at-javaone
status: publish
title: 'JUG-USA Meeting with James Gosling at #JavaOne'
wordpress_id: '200'
categories:
- javaone2009
---

![James Gosling](http://mattstine.files.wordpress.com/2009/06/dscn0912.jpg?w=220) On Wednesday morning, JUG-USA was fortunate enough to get about 45 minutes with the "Father of Java" himself, James Gosling. Each year at JavaOne, Sun distributes registration discount codes to interested JUG's. The JUG with the most registrations using their code gets this meeting. JUG-USA's founding allowed us to use a bit of strength in numbers to wrestle this crown from the bigger European JUG's for the first time that I know of.

The session was setup as an informal Q&A. I took furious notes on several of the questions that were most interesting to me, and I'll do my best to collate them into something intelligible here. _**Note that where I use quotes it's mostly to set off his answers...please don't take them as direct, but as paraphrases.**_

The first couple of questions surrounded web buzz. James sees Web 2.0 as "hype for folks who want to sell their books." When asked about the semantic web, he says his feelings are mostly positive, but maybe its an answer searching for a question. He feels that some parts of it will never perform.

When asked about Scala, I was pleased to hear him say, "I like Scala a lot. They're doing some pretty nice stuff." He talked about how functional languages lend themselves to proving properties of programs, which can then lead to automatic decomposition onto multiprocessor systems. He wishes that the Scala guys would focus more on this aspect. He also said that he fears that Scala is almost going down the Perl route of making things "too concise" and that it's still awfully hard for many programmers to "get" the functional stuff.

The next question was a rather predictable one - "James, now that you have the benefit of 14 years of experience with the Java language, is there anything you would have done differently?" Funny thing was the first thing out of his mouth was that the current AWT event model would have been one of the first things to go. Generics and closures were things he quickly mentioned that he wanted to have done, but he also said that there was quite a bit of arguing within the Java team about polishing for years vs. shipping the language. James said he felt like there was a certain window of opportunity out there that they needed to ship the language within to be successful (and it looks like they hit it). Continuing on down the list, James mentioned Eiffel-style precondtions and postconditions and type inferencing. He explained that the entire type system was designed to make C programmers comfortable, and at the time that was clearly the right decision. He state that at this point, the C/C++ comfort factor counts for nothing, and remarked that the casting syntax in Java is just "dumb." A follow-up question surrounded whether or not primitives should be a part of the language. James called this a hard problem, asserting that a "uniform type model can lead to bumpy behavior or bad performance." He said that the Scala specification gets around this with a dual object hierarchy (I need to look this up to see what he means.). He continued to talk about the definition of equality vs. identity, explaining that primitives can't have identity; in fact, they may never exist except as a figment of the compilers imagination, optimized away (my notes are fuzzy on this...when I get the recording I'll clean this up).

The next question surrounded a statement that one attendee of the recent Google I/O conference heard, that threads are "deprecated" as a concurrency model. James immediately retorted that concurrency has very many ideas around it, and that it has been an active Ph.D. thesis generator for 30+ years. He mentioned competing models such as Actors in Scala and Erlang, and Communicating Sequential Processes (CSP) process algebra from Occam. He asserted that with CSP most people typically don't "think that way very well. People with Math degrees think it's cool, but everybody else is kind of like...huh?" He stated that no data sharing between threads makes a lot of problems go away, but that often times casting problems that way can be intractable. A big quotable: "All of the different threading paradigms feel to me like whack-a-mole." He closed the concurrency question by mentioning Guy Steele's work on Fortress, where they are doing automatic decomposition from functional constructs into multiple threads of execution, making all of this issue transparent to the programmer.

Since this is getting pretty long, I'll mention that James says he knows that Larry Ellison, Oracle CEO (and Sun "suitor") has downloaded JavaFX and personally written apps with it, saying it is "pretty darn tasty!" (Again, don't take this as quoted...it's definitely paraphrased).

Other quotables: "Android - it's hard to know what to think about it. It seems like a corner hobby project, chaos looking for a place to happen."

"The exception system was DESIGNED to be a complete PITA."

All-in-all, a great time was had by all. Thanks James for your time, it was definitely fascinating hearing your thoughts on so may different areas of the Java landscape.



![JUG-USA with James Gosling](http://mattstine.files.wordpress.com/2009/06/jug-usa-gosling-grouppic-900.jpg)
