---
date: '2009-06-09 20:19:30'
layout: post
slug: javaone-2009-return-of-the-puzzlers-schlock-and-awe
status: publish
title: '#JavaOne 2009 Return of the Puzzlers: Schlock and Awe'
wordpress_id: '223'
---


For several years now, Neal Gafter (Microsoft) and Joshua Block (Google), have made a habit of presenting various incarnations of this technical session, focused on what they call "Java Puzzlers." Java Puzzlers are nothing more than short Java programs with curious behavior. It is a somewhat interactive session, with each puzzler's code listing followed by a four possible choices (A,B,C,D) to answer the question: "What does it print?" Neal and Joshua "require" that every attendee vote for their answer. After the vote, they reveal the correct answer and how to fix the problem. At the end of each problem presentation is the most important part, which is the Java programming principle that is illuminated by the puzzler. These are important tools to carry with you as you go forth a develop your Java code.

The following is a listing of the problem titles that were presented:




	
  1. Life's Persistent Questions

	
  2. Instruments of Torture

	
  3. Iterator Titillator

	
  4. Searching for the One

	
  5. Cogito Ergo Sum

	
  6. Thread Friendly

	
  7. When Words Collide



I'm going to spend my time focused on the last puzzler, as it was the one for which I was ABSOLUTELY CERTAIN that I knew the correct answer. Here's the code listing:


    
``` java    
public class PrintWords {
  public static void main(String[] args) {
    System.out.println(
      Words.FIRST + " " + Words.SECOND + " " + Words.THIRD);
  }
}

public class Words { // Compile PrintWords against this version
  public static final String FIRST  = "the";
  public static final String SECOND = null;
  public static final String THIRD  = "set";
}

public class Words { // Run against this version
  public static final String FIRST  = "physics";
  public static final String SECOND = "chemistry";
  public static final String THIRD  = "biology";
}
```


So here's the problem. We have two versions of class "Words." We compile class "PrintWords" against the first version and we then run the class against the second version. What does it print out? I used my trusty "seasoned Java programmer" knowledge to assert that the program would print "the null set." Why? Well, any seasoned Java programmer realizes that static final variables (or "constant variables" - what a curious concept) are inlined by the compiler. So it's quite obvious that "the" would be inlined wherever FIRST appears, null would be inlined wherever SECOND appears, and "set" would be inlined wherever THIRD appears.

Imagine my frustration when Josh and Neil announced that the program would in fact print "the chemistry set." As it turns out, null cannot be inlined! It is not a constant variable. So when we run our Java program, the JVM smartly picks up the non-null value of SECOND that is is able to find in the new version of Words.

Simply because this one got me, I wanted to highlight it for my readers. The "moral" of this story is that only primitives and strings can be constant, and that null is not a constant. One should use a constant field ONLY if its value will NEVER change. For final fields whose value may change, they suggest the use of an ident() method:


    
``` java 
// Returns its argument
private static String ident(String s) {
  return s;
}

// None of these fields are constant variables
public class Words {
  public static final String FIRST  = ident("the");
  public static final String SECOND = ident(null);
}
```    



The compile will not inline the call to ident.

Want more? Pick up a copy of Neal and Josh's book, _Java Puzzlers: Traps, Pitfalls, and Corner Cases!_
