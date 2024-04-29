<div align="center">
	<p>
		<img alt="Thoughtworks Logo" src="https://raw.githubusercontent.com/ThoughtWorks-DPS/static/master/thoughtworks_flamingo_wave.png?sanitize=true" width=200 />
    <br />
		<img alt="DPS Title" src="https://raw.githubusercontent.com/ThoughtWorks-DPS/static/master/EMPCPlatformStarterKitsImage.png?sanitize=true" width=350/>
	</p>
  <h3>opw - 1password connect server cli</h3>
  <a href="https://app.circleci.com/pipelines/github/ThoughtWorks-DPS/opw"><img src="https://circleci.com/gh/ThoughtWorks-DPS/opw.svg?style=shield"></a> <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/license-MIT-blue.svg"></a>
</div>
<br />
---
**DEPREACTED**: Since this tool was created, the 1password cli supports Writes to Connect or op Cloud endpoints. The opw cli is no longer supported going forward.  
---

The 1password cli (`op`) when used to interact with the 1password connect (secrets automation server), does not yet support writing/updating secret values. While this is apparently on the roadmap, it is a necessity for using 1password as a secrets store for infrastructure and service integration automation.  

This utility provides that functionality. Hopefully this is just a stop-gap measure until they release this in their own tool.  

## installation

Pre-compiled versions for various systems are available [here](https://github.com/ThoughtWorks-DPS/opw/releases).  

## quick-start

passing item field and value on the commandline. Will create the item and/or field if it does not exist, or update the Value if it does exist.  
```
$ opw write my-item my-field gagfuye62351j
```

piping value to opw cli.  
```
$ cat filename.txt | opw write my-item my-field -
``

#### contributing

* Current pipeline does not support multiple tags on the same commit.  
