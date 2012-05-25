---
date: '2008-02-05 12:27:00'
layout: post
slug: set-conditional-breakpoints-in-idea
status: publish
title: Set Conditional Breakpoints in IDEA
wordpress_id: '60'
---

So yesterday I was attempting to debug an issue in a batch processing module within one of our applications. In short, an assertion was failing deep within Hibernate as it attempted to flush the session. Using a combination of various log statements, I had isolated the problem down to a particular record that the batch process was attempting to update. (BTW: I know you shouldn't be using Hibernate for batch processing - however, we're talking about batches of at most 1000 records here, not millions!) What I really wanted to do was set a breakpoint and examine the state of the objects at runtime; however, I dreaded the thought of clicking through the breakpoint time and time again until I got to the particular record that was causing the problem. "Surely," I thought, "there must be a way to tell the debugger to only break under certain conditions."  
  
So, here's the code I wanted to examine:  

    
    <br></br>public Publication parsePublication(String inputLine)<br></br>   throws ParseException {<br></br>       Publication publication = new Publication();<br></br><br></br>       String[] fields = inputLine.split("\t");<br></br><br></br>       publication.setPublicationType(fields[0]);<br></br>       ...<br></br>       return publication;<br></br>}<br></br>

  
Essentially, I wanted to break after `inputLine.split("\t");` if and only if `fields[35]` existed and was equal to "PM:16732581." After examining IDEA's Breakpoint dialog, I noticed a section in the bottom right-hand corner that I'd never played with before:  
  
[![](http://2.bp.blogspot.com/_Vo63LRwAZbk/R6idFtXVh2I/AAAAAAAAAcw/DpuaZ5TQUCs/s320/breakpoint_dialog.jpg)](http://2.bp.blogspot.com/_Vo63LRwAZbk/R6idFtXVh2I/AAAAAAAAAcw/DpuaZ5TQUCs/s1600-h/breakpoint_dialog.jpg)  
  
As it turns out, this is exactly what I needed. If you click on the ellipsis next to the drop menu, you get a context-sensitive editor equipped with code completion:  
  
[![](http://3.bp.blogspot.com/_Vo63LRwAZbk/R6idb9XVh3I/AAAAAAAAAc4/MENimxICxyk/s400/condition_dialog2.jpg)](http://3.bp.blogspot.com/_Vo63LRwAZbk/R6idb9XVh3I/AAAAAAAAAc4/MENimxICxyk/s1600-h/condition_dialog2.jpg)  
  
Enter the desired conditions and voila! A conditional breakpoint. It worked like a charm the very first time, and I only had to inspect the breakpoint when the problematic record came up.  
  
Another nice feature of the conditional breakpoint is that if some sort of exception (such as a NullPointerException) occurs while attempting to evaluate the conditional expression, IDEA pops up a dialog informing you what happened and asking if you want to stop at the breakpoint or continue. Nice.
