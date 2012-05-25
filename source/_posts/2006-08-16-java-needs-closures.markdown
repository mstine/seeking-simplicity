---
date: '2006-08-16 09:45:00'
layout: post
slug: java-needs-closures
status: publish
title: Java needs closures
wordpress_id: '23'
categories:
- java
---

I’m sure everyone is sick of reading this same rant over and over, but I just had to add more fuel to the fire. I’m attempting to implement access privilege delegation in a JSF application - basically, users can delegate their ability to do “stuff” in our application to other users. I have a backing bean that has several methods that are called by the JSF components, returning whether or not to render that component based on security privileges. Well, I now have to make all of those methods aware of delegation! I have something like this in several methods:
    
    boolean notX = (loggedInUser.isX());<br></br>if (notX) {<br></br>  boolean result = false;<br></br>  Set delegations = loggedInUser.getDelegations();<br></br>  for (Iterator i = delegations.iterator(); i.hasNext();) {<br></br>     User delegator = (User) i.next();<br></br>     if (delegator.isX()) {<br></br>        result = true;<br></br>        break;<br></br>     }<br></br>  }<br></br>  return result;<br></br>} else {<br></br>  return true;<br></br>}<br></br>

Now, it would be nice if I could extract the contents of that if block into a new method, say “checkDelegations()”. Unfortunately, the isX() that I need to call is different for every method following this pattern. I’d like to be able to pass a function that calls isX() on the delegator into the checkDelegations() method. No dice in Java. Does anyone else have a solution to this problem? 
