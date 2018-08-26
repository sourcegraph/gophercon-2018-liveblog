# Writing Guide

Sourcegraph has liveblogged for several GopherCons so we can offer some solid tips.

This page describes how to go about creating the best possible post while the [post submission process page](post-submission-process.md) details how to get your post to us.

## Post format

We use markdown for our posts with a metadata section at the top of the template for you to fill in your name and the url you would like us to link to at the top of the blog post.

## Examples of previous liveblog posts 

Here are some past examples of liveblogs that have done well on HackerNews:

 - [Go GC: Solving the Latency Problem in Go 1.5](http://gophercon2015.tumblr.com/post/123574706480/go-gc-solving-the-latency-problem-in-go-15) (GopherCon 2015)
 - [How a complete beginner learned Go](http://gophercon2015.tumblr.com/post/123565059490/how-a-complete-beginner-learned-go-as-her-first) (GopherCon 2015)
 - [Optimizing Go programs](http://gopherconindia.tumblr.com/post/111549295932/jason-moiron-go-faster-optimizing-go-programs) (GopherCon India 2016)

 More examples can be found at [sourcegraph.com/gophercon](sourcegraph.com/gophercon).

 You'll notice some vary in length and level of detail. We trust your judgement in deciding what's right for the talk that you're covering and we've got some tips below to help guide you in writing the best possible post.

## Writing a great liveblog post

Even though it's technically a blog post, treat this more as a note taking exercise. Bullet points are your friend as they're easy to scan and avoid you having to write flowing sentences which aren't necessary.

## Be prepared (if possible)

If we can, we'll send you a draft copy of the slides for you to read beforehand. This will allow you to start your post and get a sense for the points you want to cover.

You can also extract important pictures, diagrams and code snippets from the slides and put these in your pre-talk draft.

This is, of course, not a substitute for watching the talk, but being familiar with the subject matter enables you to more easily capture new and interesting things not in the slides as you're not scrambling to both understand and capture content at the same time.

## What to capture

> For me, it's helpful to view the liveblog as a note-taking task. Pretend you're writing class notes for a friend who can't attend the talk but needs to know the key material for the test. - [Beyang Liu](https://twitter.com/beyang), Sourcegraph co-founder.

Focus on succintly summarizing the most important points and takeaways and don't worry about recording every detail mentioned. Knowing what to include and what to leave out is a bit of a learned skill but just trust your judgement and think, "What would would someone need to know in order to get value from my post?"

## Code snippets

We love code snippets in our liveblog posts and if possible, it would be great for you to type the code out instead of taking screenshots.

To enable syntax highlighting, use the following format:

```go
h := mat.Minus(mat.Product(X, theta), y)
return mat.Dot(h, h).Sum() / float64(2 * X.Rows())
```

## Image URLs

As you'll be using the GitHub issue editor to create your post (sounds weird but we'll explain later), all you have to do is drag and drop the image into the editor and GitHub will create the URL for you.

## Enjoy the experience!

Liveblogging can be crazy, frantic, but super rewarding. You get to be the one to get the word out about some really cool concepts and projects, and it's your name on the byline. 

And chances are while everyone else is checking their email or social media, you'll be actively grokking the content of the talk on a level where you can actually explain it to others. 

Consequently, you'll take away much more knowledge yourself, and you'll have lots of super interesting thoughts to share with people at dinner and drinks, making you the life of the conference after-party :).

## Next...

Now that you know how to write a post, let's get talk about [getting your post live](post-submission-process.md).
