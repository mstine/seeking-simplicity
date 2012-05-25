---
date: '2006-08-17 09:46:00'
layout: post
slug: matt-rtfm-commons-collections-has-closures
status: publish
title: Matt, RTFM! Commons Collections HAS CLOSURES.
wordpress_id: '24'
categories:
- java
---

Thanks to everyone who pointed me to [Commons Collections](http://jakarta.apache.org/commons/collections/) and its [Functor package](http://jakarta.apache.org/commons/collections/api-release/org/apache/commons/collections/functors/package-summary.html) yesterday. To me, this is just one step below having closures natively present in the language. FYI, I was able to remove all duplication from my class and reduce the LOC from 211 to 136 - in other words, 75 lines of useless code GONE.

Here is what I did. First, I defined Predicates for each of my conditions. The simplest ones called a boolean method on the User object:
    
    private final Predicate isX = new Predicate() {<br></br>  public boolean evaluate(Object object) {<br></br>     return ((User) object).isX();<br></br>  }<br></br>}<br></br>

Only slightly more complicated ones checked to see if a given Collection was empty on the User object:
    
    private final Predicate isX = new Predicate() {<br></br>  public boolean evaluate(Object object) {<br></br>     return CollectionUtils.isNotEmpty(((User) object).getItems());<br></br>  }<br></br>}<br></br>

Next, I defined a method that would check the delegations for the User to see if any of them were an X:
    
    boolean checkDelegations(User user, Predicate checkPredicate) {<br></br>  return CollectionUtils.exists(user.getDelegations(), checkPredicate);<br></br>}<br></br>

Finally, I implemented the security methods:
    
    public boolean canDoThis() {<br></br>  return isX.evaluate(loggedInUser) || checkDelegations(loggedInUser, isX);<br></br>}<br></br>

Maybe it isn’t the most elegant or simplest of solutions, but it sure is a lot better than what I posted yesterday!

P.S. Since this is a Christian blog, I must remind you that RTFM stands for Read The **FINE** Manual! 
