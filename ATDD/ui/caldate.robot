*** Settings ***
Library           SeleniumLibrary


*** Variables ***
${SERVER}         http://localhost:8888/caldate.html
${BROWSER}        Chrome


*** Test Cases ***
Test calculate
    Open web
    Input fromdate and todate
    press calculate button
    display result


*** Keywords ***
Open web
    Open Browser    ${SERVER}    ${BROWSER}

Input fromdate and todate
    Input Text     id=startDD    4
    Input Text     id=startMM    1
    Input Text     id=startYYYY    2018
    Input Text     id=endDD    4
    Input Text     id=endMM    7
    Input Text     id=endYYYY    2018    

press calculate button
    Click Button    id=btnCalc
    
display result    
    Wait Until Page Contains    Thursday, 4 January 2018
    Wait Until Page Contains    Wednesday, 4 July 2018
    Wait Until Page Contains    Result: 182 days
    Wait Until Page Contains    15,724,800 seconds
    Wait Until Page Contains    262,080 minutes
    Wait Until Page Contains    4368 hours
    Wait Until Page Contains    26 weeks
    Wait Until Page Contains    49.86% of 2018
    