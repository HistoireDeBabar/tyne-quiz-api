### Quiz

###### Get

    /quiz?=id
    

  **Response**

    {
      "id": "test",
      "questions": [
        {
          "id": "a",
          "question": "Infamous Newcastle United number 9 and the English Premier Leagues all time top goal scorer",
          "Answers": [
            {
              "id": "1",
              "QuestionId": "a",
              "Correct": true,
              "AnswerCount": 2,
              "Answer": "Alan Shearer"
            }
          ]
        }
      ]
    }


###### Post

    /answer


**Request**


    {
        "id": "test",
        "answers": [
            {
                "questionId": "a",
                "id": "1"
            }
        ]
    }

