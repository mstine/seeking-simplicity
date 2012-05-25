---
date: '2010-06-03 17:00:35'
layout: post
slug: dont-build-software-thats-too-smart
status: publish
title: Don't build software that's TOO smart!
wordpress_id: '355'
categories:
- agile
- development
---

I had an extremely successful meeting with one of our clients yesterday. We were discussing how we wanted to go about migrating her laboratory from its current system (one that we built several years ago) to our new lab management platform. At some point during the discussion I made the statement, "We tried to make the previous system too smart! We're not repeating that mistake this time." Of course, she was in complete agreement with that principle. I've had similar interactions with our other clients that are making migrations (rather than encountering our system for the first time on this new version), and although I've never explicitly stated the principle that way, similar sentiments have abounded.

What is too smart software? In our case, it was a system that attempted to encapsulate every single rule of "business" process that occurred within a given laboratory. Many times statements were flung around like "will it ALWAYS happen this way," "what should we do if this happens?" etc., etc., etc. We tried to cover every single possibility, and we did an excellent job of preventing users from ever breaking their own rules. What we didn't realize (and we're not unique - this problem is RAMPANT) is that the rules CHANGE. Rules come, rules go. Sometimes the rule remains, but there are a few exceptional cases that must be dealt with. Our system simply couldn't deal with a world that worked this way - and thus, our system was completely unfit for the real world.

We set out with a different mission this time. If there's one overriding characteristic of SRM (Shared Resource Management) 2.0, it's the explicit assumption that the world will change continually. We don't attempt to tell you how you must use this system. We capture your data, we invoice for your services, we run your reports - but YOU, the user gets to decide how you'll interact with it. If your workflow changes, we change with you. Now the devil is in the details. It's taken roughly 20-30 man years worth of effort to build a system like this, and it hasn't been easy. But in the end, we're finding that those years were much better spent ENABLING our users rather than PREVENTING our users from getting things done. 

I'm not sure that I've gotten my point across in this brief diatribe, so I'll attempt to sum it up here. If you're developing a system, figure out the 2 or 3 things that will make your users' lives AWESOME, and do those 2 or 3 things extremely well. Don't do the rest AT ALL. Don't build a system that attempts to be smarter than the knowledge expert using it - it's a means to your user's end, not an end in itself. 
