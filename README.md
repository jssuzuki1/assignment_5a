# Webscraping Wikipedia

GitHub Repo URL: https://github.com/jssuzuki1/assignment_5a

Our main objective is to create a program that will scrape the following wikipedia articles using the Colly package:
		- "https://en.wikipedia.org/wiki/Robotics",
		- "https://en.wikipedia.org/wiki/Robot",
		- "https://en.wikipedia.org/wiki/Reinforcement_learning",
		- "https://en.wikipedia.org/wiki/Robot_Operating_System",
		- "https://en.wikipedia.org/wiki/Intelligent_agent",
		- "https://en.wikipedia.org/wiki/Software_agent",
		- "https://en.wikipedia.org/wiki/Robotic_process_automation",
		- "https://en.wikipedia.org/wiki/Chatbot",
		- "https://en.wikipedia.org/wiki/Applications_of_artificial_intelligence",
		- "https://en.wikipedia.org/wiki/Android_(robot)",

This folder includes everything needed to execute that program. 

## Files Overview and Descriptions

- main.exe : the deliverable for this objective. It is an executable load module that will scrape a set of wikipedia articles. To execute this program, simply double-click the application or go into terminal, change the directory to the folder that contains the main.exe file, and run the command '.\main.exe'
- scraped_wikis.json : main.exe produces this file. This file structures the scraped data in the following form: {url, text}
- main.go : this is the code that main.exe runs. 
- main_test.go : this is the corresponding test file for main.go. It has very similar code to the original main.go file but essentially loops through main.go 100 times as a benchmark test.
- scraped_wikis_test.json : main_test.go produces this file. This file structures the scraped data in the following form: {url, text}
- go.mod, go.sum, go.work.sum: all of these files include the dependences that main.exe needs in order to run. Most notably, all of this contains the necessary packages to run Colly. 

## Evaluation of Performance

main_test.go looped through 100 iterations of the code in main.go. The entire process took 35.25 seconds, or 0.3525 seconds per iteration on average. 

For collecting all of the text of 10 wikipedia articles, this is very quick. 

So, management! Listen up. GoLang doesn't just leave the house like a half-awake burned out quantitative analyst forced to copy and paste text into an excel spreadsheet. It Go Fast. 
