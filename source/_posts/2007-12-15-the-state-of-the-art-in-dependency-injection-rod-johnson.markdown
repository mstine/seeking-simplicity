---
date: '2007-12-15 10:06:00'
layout: post
slug: the-state-of-the-art-in-dependency-injection-rod-johnson
status: publish
title: 'The State of the Art in Dependency Injection: Rod Johnson'
wordpress_id: '52'
categories:
- java
- spring
- thespringexperience2007
---

This was a truly fascinating talk. If you ever wanted to learn the entire history and landscape of dependency injection (DI), this was your opportunity. I really didn't realize how deep of a topic DI really is.  
  
According to Rod, DI had its beginnings in 2002, in the Interface 21 Framework that was born from his seminal work, [Expert One-on-One J2EE Design and Development](http://www.amazon.com/Expert-One-One-Design-Development/dp/1861007841). In the beginning, DI was done solely through setter injection (SI), with external metadata (usually in XML). 2003 gave us Spring 0.9, which had the same DI model, but added FactoryBeans for indirection and proxy-based AOP (enabling among other things, declarative transaction management). With Spring, DI was always just one enabler of a complete enterprise solution. Contrast that with [PicoContainer](http://www.picocontainer.org/), also arriving on the scene in 2003, which was an ultra-lightweight DI-only container. PicoContainer brought us several innovations, including constructor injection, automatic resolution by type, and an attempt to dispense with external configuration.  
  
Rod then discussed the Pros and Cons of Constructor Injection (CI). On the Pro side, we see that CI is great for immutable objects, can be used with existing code, enforces that the necessary dependencies are provided at object construction, and allows developers to dispense with methods like afterPropertiesSet(). On the Con side, there are no default arguments in Java, which forces us to ALWAYS provide all of the necessary dependencies, constructor overriding can be somewhat messy, and constructor argument names are not usually available through reflection, which forces us to depend on argument order for dependency resolution.  
  
Spring 1.0, which arrived in late 2003, incorporated some of PicoContainer's innovation: CI, "Autowiring" by type and by name, as well as the ability to mix CI and SI. Here Rod reflected on Spring's pragmatic philosophy, whereby they are more than willing to incorporate the good ideas that others had first.  
  
Spring 1.1 brought us method injection, a unique Spring feature, which solved the problem of long-lived managed objects which depend on short-lived objects. It also brought us Factory Methods. You now had the ability to create objects through invoking static factory methods or instance methods on other beans.  
  
With the finalization of the [EJB 3.0 standard](http://java.sun.com/products/ejb/docs.html) in May 2006, several new ideas became available in the world of DI: annotation-based DI, annotations for callback methods (e.g. @PostConstruct), scanning for annotations without any external configuration needed for a deployment unit, and field injection (FI). These EJB standard brought with it a new philosophy - that annotations are a superior model and that externalization should be a second-class citizen. The Pros of this standard? One big one - no configuration is necessary for simple cases. The Cons? Rod listed several: resources could only be injected from JNDI, no CI, limited control over object lifecycle, unsuitability for fine-grained object graphs, effectively usable ONLY with annotations, as the XML configuration option is so verbose, and it isn't just DI - it brings along with it enforcement of many obsolete EJB concepts (such as bean instance pooling).  
  
Spring 2.0 arrived in late 2005 with its own share of innovations: integration with [AspectJ](http://www.eclipse.org/aspectj/), the ability to inject anything with an AspectJ aspect via @Configurable, namespaces to provide an XML-based DSL for external configuration, the @Required annotation, and pluggable scopes to manage the object lifecycle. About that same time, JBoss Seam introduced many strange ideas (that I didn't quite grasp from Rod's explanation): bijection, outjection, and subversion of control.  
  
Some additional players on the field include Spring Java Config, which is a DSL for configuration in Java, [Google Guice 1.0](http://code.google.com/p/google-guice/), and the [JSR-299 Web Beans spec](http://jcp.org/en/jsr/detail?id=299), which Rod described as "blogware" in that there is currently no available implementation.  
  
Finally we got to Spring 2.5, released last month. Rod didn't have his slides printed or available on the web (argh!), so about this time my hand was really hurting from taking furious notes, and I simply gave up. Fortunately, most of what Rod covered was redundant from some of the other sessions, so I'll have or will blog about it later. About the only thing I really don't have good notes on is Spring Java Config, which he discussed pretty extensively. If you want to learn more about it, check out [http://www.springframework.org/javaconfig](http://www.springframework.org/javaconfig).
