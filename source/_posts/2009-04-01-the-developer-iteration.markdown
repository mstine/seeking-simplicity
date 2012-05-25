---
date: '2009-04-01 16:56:59'
layout: post
slug: the-developer-iteration
status: publish
title: The Developer Iteration
wordpress_id: '96'
categories:
- agile
- CodeProject
---

I recently gave a 90 minute talk at work covering quite a bit of ground related to Agile Software Development. If I had to sum up agile development in one word, it would be FEEDBACK. Venkat Subramaniam and Andy Hunt give one of the best definitions of agile I've ever seen in their book, _Practices of an Agile Developer_:


> _Agile development uses feedback to make constant adjustments in a highly collaborative environment._


Here's my graphical representation of this quote:


![Agile Feedback Loop](http://mattstine.files.wordpress.com/2009/04/aws_fb_1012.jpg?w=300)



I decomposed this into an "ecosystem" of multiple feedback loops:


![Developer to Code](http://mattstine.files.wordpress.com/2009/04/aws_fb_2013.jpg?w=300)




![Developer to Developer](http://mattstine.files.wordpress.com/2009/04/aws_fb_3025.jpg?w=300)




![Team to Client](http://mattstine.files.wordpress.com/2009/04/aws_fb_4033.jpg?w=300)



What I'd like to focus on in this post is what I consider to be the heart of the "Developer to Code" feedback loop: **The Developer Iteration**.

You see, iterations aren't just for teams. Indeed they are also for developers. So many times I have witnessed individuals struggling to participate fully in an agile project because when they return to their desks, for all intents and purposes, they are pretending to be a waterfall. They spend days analyzing a feature, code for days without sharing any code, test the happy path toward the end of the iteration, and attempt a big bang integration on the last day. I'll be the first to admit that I myself tried to code this way and had a couple of realizations after a short time:



	
  * Something felt completely unnatural about it, especially in the context of what we were trying to do as a team...

	
  * I was awfully miserable, especially at the end of the iteration.


So, what does the developer iteration look like? It is essentially a microcosm of events that occur during a regular team iteration:

	
  1. Look at your feature or task, and plan out what you have to do.

	
  2. Break the work into small chunks, where each chunk leaves the system in a stable state, yet that much closer to the feature or task being complete.

	
  3. Develop each chunk (preferably using Test-Driven Development), unit and integration testing your work along the way.

	
  4. Check-in your code EVERY TIME you again arrive at a stable state.


I'll give my definition of stability: A system with NO broken tests. How do we maintain stability? Here are several "smaller" steps that should immediately proceed your check-in step:

	
  1. Run the tests on your development machine, fixing all broken ones until you reach 100% passing with the code you have.

	
  2. Check out the latest code from your version control system, thereby integrating your team's changes with your own.

	
  3. Run the tests AGAIN, fixing all broken ones until you reach 100% passing with the integrated code. At this point, your system is stable locally.

	
  4. Check-in the code! Now the entire team has a stable, fully-integrated system.


So, what's the point of all this? Look at how much feedback you're giving yourself along the way! Each time you stop and run the tests, you're getting feedback about the stability of your code. Not only that, but your constant striving for 100% stability greatly enhances your ability to refactor your code to keep it clean. You're able to constantly refine the code, creating smaller methods, more cohesive classes, and an overall simpler and better design combined with the confidence that your system is functionally equivalent to what you started with prior to refactoring.

Yet another benefit is an enhanced ability to timebox yourself. Working this way enhances your ability to plan to be "done" by the end of the day, as you're never very far from your next stable state. Go home with a stable system, and find a stable system when you get back in the morning. Doesn't that feel great? Even if you can't reach a stable state by the time you head home, there's always the throwaway option. Just revert your changes and go home with stable code. Try again with a fresh mind the next morning. Again, if you're working the developer iteration, you haven't written that much code since your last stable state anyway. :-)

Try it out. I'm not a salesman. I'm a satisfied customer.
