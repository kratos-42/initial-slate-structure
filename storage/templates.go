package storage

var skeleton_template string = `---
title: API Reference

language_tabs: # must be one of https://git.io/vQNgJ
  - shell: cURL

toc_footers:
  - <a href='https://github.com/ReiKratos/initial-slate-structure'>Initial Slate Structure</a>

search: true
---

# Introduction

`

var query_string_options_template string = `### Query String Options

It is possible to:

`

var example_template string = "> Example request:\n\n" +
                              "```shell\n(Insert example here)\n``` \n\n" +
                              "> Example Response:\n\n" +
                              "```json\n(Insert response here)\n```\n\n"

var include_table_template string = `* Include related data:

Option | Required | Example
------ | -------- | -------

<br>

`

var sort_table_template string = `* Sort results:

Option | Required | Example
------ | -------- | -------

<br>

`

var filter_table_template string = `* Filter results:

Option | Operators | Required | Example
------ | --------- | -------- | -------

<br>

`

var paginate_table_template string = `* choose how the results are paginated by passing the number of:
  * results *per* page
  * the page to retrieve

Option    | Required | Example
--------- | -------- | ------------
` +
"size      | no       | `page[size]=2`\n" +
"number    | no       | `page[number]=3`" + `

<br>

`

var after_middlewares string = "**Example query string:** <br> " +
                               "`?{add your query string example here}`\n\n" +
`### Headers

Key                 |  Value            | Required
------------------- | ----------------- | --------
Accept-Language     |    {locale}       | no

### Path Variables

Variable            |  Required
------------------- | ------------

Status Codes:

<aside class="success">
{Add the success result here}
</aside>

<aside class="warning">
{Add the error code here. If more than one, use more 'aside' tags with warning class}
</aside>

`
