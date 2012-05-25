---
date: '2009-04-25 00:14:40'
layout: post
slug: grails-prototype-scriptaculous-persistent-grid-sorting-goodness
status: publish
title: Grails, Prototype, Script.aculo.us - Persistent Grid Sorting Goodness
wordpress_id: '139'
categories:
- ajax
- CodeProject
- grails
- groovy
- prototype
- scriptaculous
---

Ever wanted to do drag-n-drop sorting of a grid of images on a page and persist it? Here's my solution using Grails, Prototype, and Script.aculo.us.

Basically what prompted this was the need for my wife to be able to sort the various product images that she had on a screen at any given time in any way that she pleased, and it had to be easy to work with. What follows is by no means a complete solution to this problem, but it represents where I am in the development process and may be useful to you, my hapless reader.

Here's some GSP code which basically lays out a grid of product images, 3 wide by _n_ deep:


    
    
    <div id="productThumbContainer">
      <g:set var="rowIndex" value="${0}"/>
    <g:each in="${products}" var="product" status="index">
      <g:if test="${index % 3 == 0}">
        <div id="productRow${rowIndex}" class="span-20 last product-row">
      </g:if>
      <div id="product_${product.id}" class="span-6 product <g:if test="${(index % 3 == 2) || ((products.size() - index) == 1)}">last</g:if><g:else>append-1</g:else>">
        <img src="${resource(dir: grailsApplication.config.store.productImages.webPath, file: product?.thumbnailImage?.name)}" width="230" class="productImage">
        <h3><g:link controller="product" action="show" id="${product.id}">${product.name}</g:link></h3>
      </div>
      <g:if test="${(index % 3 == 2) || ((products.size() - index) == 1)}">
        </div>
        <g:set var="rowIndex" value="${rowIndex + 1}"/>
      </g:if>
    </g:each>
    </div>
    



Now here's where the magic happens:


    
    
    document.observe('dom:loaded', function() {
          var productRows = $$('.product-row');
    
          var options = {
            constraint: false,
            overlap: 'horizontal',
            containment: productRows,
            dropOnEmpty: true,
            tag: 'div',
            onUpdate: updateRows
          };
    
          productRows.each(function(item) {
            Sortable.create(item, options);
          });
    
          $('persistOrderingButton').observe('click', function(event) {
              var sortString = '';
              productRows.each(function(row) {
                  sortString += '&';
                  sortString += Sortable.serialize(row);
              });
              <g:remoteFunction action="sortProducts" params="sortString" update="ajaxMessage" onSuccess="\$('ajaxMessage').show()"/>
          });
    });
    



What we've got here is, failure to communicate...oops, wrong synapse there...what we've got here is a Prototype selector that grabs everything with the class ".product-row." It then takes these and creates a Scriptaculous Sortable for each of them, providing the object-literal "options." Notice the "containment" option which allows you to drag products back and forth between the various rows.

Delving deeper into the magic is the callback function "updateRows." This guy is my favorite Javascript creation in quite some time:


    
    
    function updateRows(list) {
          var children = list.childElements();
    
          if (children.size() < 3) {
    
            //If I'm the last row, who cares!
            if (list.next() != null) {
              var prevRow = list.previous();
    
              if (prevRow != null) {
                var lastChild = prevRow.childElements()[prevRow.childElements().size() - 1].remove();
                list.insert({top:lastChild});
                updateRows(prevRow);
              } else {
                var lastRow = list.up().childElements()[list.up().childElements().size() - 1];
                var lastChild = lastRow.childElements()[lastRow.childElements().size() - 1].remove();
                list.insert({top:lastChild});
                updateRows(lastRow);
              }
            }
          } else if (children.size() == 3) {
            //Do nothing...gets me out of the recursion I hope!
          } else {
            var nextRow = list.next();
            var lastChild = children[children.size() - 1].remove();
    
            if (nextRow != null) {
              nextRow.insert({top:lastChild});
              updateRows(nextRow);
            } else {
              var topRow = list.up().childElements()[0];
              topRow.insert({top:lastChild});
              updateRows(topRow);
            }
          }
    
          var i = 0;
          Sortable.sequence(list).each(function(item) {
            var productId = 'product_' + item;
            if (i < 2) {
              $(productId).removeClassName('last');
              $(productId).removeClassName('append-1');
              $(productId).addClassName('append-1');
            } else {
              $(productId).removeClassName('last');
              $(productId).removeClassName('append-1');
              $(productId).addClassName('last');
            }
            i++;
          });
    }
    



This function is organized around the fact that the only valid state for a grid of n-rows will be n-1 rows of 3 products, followed by one row of 1 <= numProducts <= 3. In most cases, if you drag a product from one row to another, you are violating that rule by creating a row with 2 products and a row with 4 products. This function solves the problem by recursively shifting the products down until reaching a stable state again.

There's a bit of noise there at the bottom of the function. I'm using Blueprint CSS to do the layout for this application and I need to shift the various Blueprint classes around once everything is settled.

Finally, here's the persistence of the data when we click save:


    
    
    def sortProducts = {
        TreeMap rowMap = new TreeMap()
    
        params.each {key, value ->
          def matcher = key =~ /productRow(.*)\[\]/
          if (matcher.matches()) {
            def rowId = matcher[0][1]
            rowMap[rowId] = value
          }
        }
    
        def productIds = []
        rowMap.values().each { row ->
          row.each {
            productIds << it.toLong()
          }
        }
    
        shoppingService.saveSortOrder(productIds)
    
        render("Product sort order saved!")
    }
    



and the logic from shoppingService.saveSortOrder():


    
    
    def saveSortOrder(def productIds) {
        def productsToSort = Product.findAllByIdInList(productIds)
    
        def productMap = [:]
        def sortIndexList = []
    
        productsToSort.each {
          productMap[it.id] = it
          sortIndexList << it.sortIndex
        }
    
        sortIndexList.sort()
        sortIndexList = sortIndexList.reverse()
    
        productIds.each {
          productMap[it].sortIndex = sortIndexList.pop()
        }
    
        productsToSort.each {
          it.save()
        }
    }
    



You might wonder why this is so complex. What I haven't fully described is the way products are organized into a hierarchy of various categories. When you're sorting a screen, you're only sorting a subset of the products that are in that particular category. However, the sort order is maintained across the entire product set in the database. So, this logic basically just shifts around the existing sort indicies, placing them in their new relative order.

Anyway, I don't know how generally applicable this code is, but I had fun writing it and I hope you can get some use out of it. Cheers!
