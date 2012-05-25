---
date: '2007-05-07 15:09:00'
layout: post
slug: java-university-developing-enterprise-applications-with-the-spring-framework
status: publish
title: 'Java University: Developing Enterprise Applications with the Spring Framework'
wordpress_id: '28'
categories:
- javauniversity
---

This was the first session that I attended at Java University. You could take either a full-day course or two half-day courses - I elected to do the latter. Keith Donald, project lead for Spring WebFlow, presented. I have quite a bit of Spring experience already through AppFuse, but I was hoping to get a more complete view of Spring as most of what I have learned is from tinkering with existing systems. Keith promised that we'd build a system from the ground up, so it sounded like I'd get what I bargained for.  
  
And build from the ground up we did. I'd say we were almost halfway through the entire session before we even TOUCHED Spring. Up until that point all we had done was do test driven development on all of the business entities and logic. It was very interesting to me that Keith's coding methodology is quite different from what I typically find myself doing. He implemented almost all of the business logic of the application without any regard to the eventual supporting technologies (i.e. ORM frameworks, web frameworks, etc.). It seems like I almost always bind my application to a set of technologies (say Hibernate/Spring/JSF) before I ever write a line of business logic. What did he get from this? Well, his application had fully implemented and tested business logic that would work with nearly any of the available supporting frameworks. Good stuff.  
  
Another thing that was interesting to me is his disdain for packaging classes according to application layer. I can't think of an application that I've built recently where the DAO's weren't in a dao package, the "Manager's" weren't in a manger/service package, and the web classes weren't in a web package. Keith totally changed all of this. Our two main entities in this application were account and restaurant. Anything that had to do with an account (be it a DAO, service bean, web form, etc.) went in the account package. The same went for the restaurant package. I'm not sure if I like this or not, but he seemed convinced of the value.  
  
Overall it was a good session. He incorporated A LOT of live coding, and almost all of it was done from scracth - he had very little pre-made skeleton code. The drawback to that was that he only made it through 6 of the 8 sections of his slides. I wish he had backed down on one or the other.
