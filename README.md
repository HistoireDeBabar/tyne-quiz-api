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
          "answers": [
            {
              "id": "1",
              "questionId": "a",
              "correct": true,
              "answerCount": 2,
              "answer": "Alan Shearer"
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

