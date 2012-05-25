---
date: '2009-06-03 13:16:26'
layout: post
slug: javaone-2009-script-bowl
status: publish
title: '#JavaOne 2009 Script Bowl'
wordpress_id: '193'
categories:
- javaone2009
---

The JavaOne 2009 Script Bowl was quite a delight to watch. The players were Jython (Frank Wierzbecki), Groovy (Guillaume LaForge), Clojure (Rich Hickey), Scala (Dick Wall), and JRuby (Thomas Enebo). The event was divided into two rounds:




	
  1. Demonstrating Language Features

	
  2. Demonstrating Community Contributions



Each of the players focused on different angles during the language feature round. The Jython focus was squarely on the readability of the language. Frank's key quote was Python looks an awful lot like "executable pseudocode." The Groovy focus was a subset of Guillaume's "What's new in Groovy 1.6" talk, and highlighted the new compile-time metaprogramming/AST transformations. Rich Hickey directly attacked concurrency using Clojure, demonstrating a multi-threaded "Life" demo that highlighted the power of software transactional memory.

I really enjoyed Dick's angle for Scala - demonstrating that most of the Java small language changes contained in Project Coin are already present in the Scala language and very useable! Thomas Enebo went straight for the WOW factor by showing how easily a Ruby DSL can be wrapped around existing Java libraries with JRuby to create a very compelling 3D game.

Next came the community round. This was actually the most impressive part of the whole session. Frank kicked off the Jython community focus by highlighting an [interactive "art IDE" called Field](http://openendedgroup.com/field/wiki/OverviewBanners2) that drives [the Processing language](http://processing.org/). The chose Jython as the scripting language for the environment. Very neat.

The Groovy demo was something to behold. Guillaume started off by demonstrating SwingPad, an exploratory IDE for building Swing applications with Groovy. But the real treat was a very impressive Twitter+NASA Worldwind mashup built very concisely using Griffon. I really need to give Griffon a try.

Rich Hickey demonstrated the power of Clojure to build a data-driven, functional, template based website. It took an existing HTML template (with 0 code) and transformed it using "css-like" selectors provided by a [enlive](http://github.com/cgrand/enlive/tree/master). Seemed a lot like what you can do with XSLT, but much cleaner.

Dick demonstrated scouchdb, a Scala library for interacting with CouchDB.

Thomas' focus was on just how big the Ruby community is, from developers, to conferences, to books, to IDE's, to libraries.

At the end of the session, the winner was decided by crowd noise, with Groovy winning the day and Scala coming in second. My take home from this session is that really all of these languages bring something very interesting and powerful to the table, and I hope all of them continue to succeed. Great job everyone!
