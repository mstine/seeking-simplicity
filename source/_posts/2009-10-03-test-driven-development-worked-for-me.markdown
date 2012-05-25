---
date: '2009-10-03 17:02:36'
layout: post
slug: test-driven-development-worked-for-me
status: publish
title: Test Driven Development Worked for Me!
wordpress_id: '254'
categories:
- agile
---

Our team got to spend a few days with Jared Richardson this week, talking tech leadership, agile, and automated testing. At some point during the opening day's discussion, I related the story of how I initially got into Test Driven Development (TDD) and how it seriously ramped up my productivity and decreased my defect count.

In early 2003, we were slowly dragging through development of the first version of our Shared Resource Management system. At the time our development model looked a little something like this:

1) Code until you run into an unimplemented dependency;
2) Repeat until one of the dependencies is checked in;
3) Big Bang Integration!
4) Debug, debug, debug...
5) Repeat until something pseudo-demoable is ready.

As you might imagine, this wasn't a very productive way to work, but unfortunately it was all most of us knew. At the time I thought we really needed to improve on our testing so that we could eliminate a lot of the defects that at the time I didn't realize were being caused by our development model. I asked my boss if I could start a QA effort and not too much later found the XP/Agile Universe conference that was held in August 2003. It was my hope that we could learn about how to properly test/QA our software there as there seemed to be a lot of talks about this "TDD" thing. I got approval to attend, and while down there I was immersed in this new culture of Extreme Programming (XP) and Agile.

At one point I ventured into a room that was setup like what many of us would call a SCRUM/Team room today, where a group of folks were all coding away on a system and making its "Fit" tests pass. I quickly paired up with a gentleman and we went to work. At some point Brian Marick came over and asked what we were working on. I fumbled through a description of the requirement and he then asked, "Well, where's your test?" I said that we didn't have one, and he then said something that at the time was very profound to me, "Well, you can't write any code until you have a failing test!"

I spent most of my time waiting in the airport to return to Memphis reading Uncle Bob's "Agile Software Development: Principles, Patterns, and Practices." Needless to say, I was fired up.

When I returned to work in the following weeks, I essentially tried two different tactics in parallel to share what I learned with my team.

First, I practiced what Jared calls "2x4 Development," or, beating your team in the shins with a 2x4 until they start doing TDD, pair programming, continuous integration, etc. I even delivered my own version of the agile manifesto to the entire department, and the most profound comments I got out of that were "I can't ever imagine us paying two developers to sit at the same computer and work on the same code," and "Well, who's going to test the test code?!?!"

Needless to say, this didn't work very well. My second tactic was much more guerrilla in nature. I told the guy that I was working with that I didn't care what he did with the rest of his time, as long as he checked in his interfaces before he started to code. Now I had something to mock! I then proceeded to TDD all of my code. I advertised my productivity in terms of test cases passing with a little counter on my whiteboard. Development was fun again, as I could get instant feedback on how my code was working and I knew that it was interacting properly with its dependencies. Sure enough, when those big bang integrations happened, my code rarely seemed to be the source of the defects, and I quickly found myself getting pulled into other parts of the code to help develop since I had already gotten my typical assignment finished!

We delivered the section of the system that I developed using TDD much faster than any part that I had worked on up to that point (no hard numbers, but I have a pretty good memory of that time), and in addition, I wrote the CLEANEST code I had ever written up to that point.

So, while Jared wouldn't recommend you try TDD if the rest of your team isn't on board - you're never going to get someone who doesn't care about tests to fix one that they broke - this guerrilla tactic worked well for me and got me some recognition as a guy who could write good clean code fast! If I could go back to that point, I think I would have tried [Defect Driven Testing](http://www.jaredrichardson.net/blog/2005/11/03/) first or in addition to my GTDD.

So, how are you getting your team test infected?
