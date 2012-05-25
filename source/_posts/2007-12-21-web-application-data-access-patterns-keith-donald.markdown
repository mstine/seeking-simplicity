---
date: '2007-12-21 15:48:00'
layout: post
slug: web-application-data-access-patterns-keith-donald
status: publish
title: 'Web Application Data Access Patterns: Keith Donald'
wordpress_id: '54'
---

Well, after a short hiatus of getting back to "real" work, I'm going to attempt to squeeze in another Spring Experience 2007 entry. This particular session was my first of Friday morning. I've always enjoyed Keith Donald as a speaker, having first heard him at [TheServerSide.com's Java In Action conference in 2005](http://www.theserverside.com/tt/articles/content/JIApresentations/Donald.pdf), and again during his [Java University session at JavaOne 2007](http://mattatjavaone2007.blogspot.com/2007/05/java-university-developing-enterprise.html).  
  
Keith began with a discussion of why we need patterns. In his opinion, patterns simply provide a way of thinking about a particular approach to a problem. One could liken them to blueprints. They typically have a name, a description of the problem being solved, a description of the solution approach, and finally a selection of implementation guidelines and considerations. Pretty nice summation.  
  
The particular patterns that Keith planned to discuss surrounded the different data access patterns typically found in web applications.  
  
Keith began by discussing the "Unit of Work" pattern, which he defined as a unit of interaction with a shared resource that executes independently of other units. These interactions exhibit [ACID properties](http://en.wikipedia.org/wiki/ACID), and a database transaction provides a typical  example. In the context of an ORM such as Hibernate, a Unit of Work tracks changes to the object graph during a transaction and then synchronizes those changes with the database when the transaction is complete. Keith followed this discussion with an implementation example from his "Reward Network" sample application.    
  
Keith followed this discussion with two of what I'll call "subpatterns" of the Unit of Work pattern: Pessimistic and Optimistic Locking. With pessimistic locking, you prevent other units from accessing your data during a given unit of work, and then release the "lock" as soon as possible. This locking is performed to prevent a [race condition](http://en.wikipedia.org/wiki/Race_condition), resulting in an inconsistent state in your model. Keith briefly illustrated implementation of pessimistic locking with Hibernate.  
  
It's always interesting when someone asks a slightly off-topic question which ends up in a take home gem of information. At this point, someone asked if the implementations that Keith was showing showcased Spring's Hibernate support. To this Keith surprised me by answering (and I paraphrase), "Hibernate's native API's are much more mature and powerful now, so we recommend you use those rather than Spring's Hibernate API wrapper. Why add another dependency?" He then proceeded to recommend Rossen Stoyanchev's Saturday morning session on "Working with Hibernate in a Spring 2.5 Environment" (blog entry to come later!). One of the better quotes of the conference came during this aside as well:  


>   
_Whenever you can eliminate a dependency, you should._  
- Keith Donald  


  
Well, back to the presentation. Keith then moved into a discussion of Optimistic Locking. Here, within a business transaction that executes across a series of system transactions, you use optimistic locking to prevent corruption of shared data. An examples is where several users can edit account information concurrently, but the frequency of concurrent access is fairly low. Data integrity is still at risk, so you use optimistic locking to handle those infrequent cases. Keith them demonstrated the use of [Hibernate versioning to implement optimistic locking](http://www.hibernate.org/hib_docs/reference/en/html/transactions.html#transactions-optimistic).  
  
Keith next moved into a discussion of the controversial "Open Session in View" pattern. This is a specialized technical pattern where by the Hibernate session is kept open for the duration of a given web request. It is generally used as a "quick fix" to prevent lazy loading exceptions in the view tier. This is generally considered a bad thing, as you end up with a great deal of "N+1 Select" situations while rendering the view, resulting in way too much SQL executed against the database. Typically, the need to use this pattern reveals a flaw in the design of your model or queries. For example, since you generally know how far in the object graph you'll need to go for a specific view, you should load those relationships eagerly for that use case. You may also consider whether or not a given object association is appropriate to begin with.  
  
The final major pattern that Keith discussed was the "Conversational Persistence Context" pattern that is present in the upcoming Spring Web Flow 2.0 release. It is essentially a more flexible variation of the Open Session in View pattern, where by you can have an open session spanning multiple requests during a long-running user interaction context, or "conversation." To be perfectly honest, I still don't quite "get" this pattern beyond that, probably because I don't know much about Web Flow or it's cousins (e.g. Seam, Web Beans JSR-299).  
  
Keith closed with a very brief description of Master Detail, where you show an overview of objects in a collection and then allow drilling down into object details, and Paging, where you allow users to easily browse through a list of items of interest. I raised a point of contention with Keith at this point. His example implementation showed delegation of the paging concern to the DAO layer, rather than implementing that concern in the View layer. When I asked him where he thought it belonged, he thought it was obvious that it belonged in the DAO layer. I then let on that we had this debate in our team and decided the the concern belonged in the View layer. Why should I run the same query multiple times just to page through the data? Why not load the entire data set and then let the View tier decide how to page through it? Keith gracefully acknowledged my point.  
  
All in all, this was a great session, if for no other reason than it reminded me of some best practices in this area and let me in on the little secret about using native Hibernate API's rather than Spring's Hibernate template. Nice job, Keith.
