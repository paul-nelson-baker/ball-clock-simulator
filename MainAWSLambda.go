// +build aws

package main

import (
	. "github.com/paul-nelson-baker/ball-clock-simulator/support"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"encoding/json"
)

type BallSimulationRequest struct {
	BallCount      int  `json:"BallCount"`
	IterationCount *int `json:"IterationCount,omitempty"`
}

type BallClockSimulationResponse struct {
	Days      *int      `json:"DaysToRecycle,omitempty"`
	Seconds   *float64  `json:"SecondsToCalculate,omitempty"`
	BallClock BallClock `json:"ClockState"`
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Get the request and validate it
	var ballSimulationRequest BallSimulationRequest
	if err := json.Unmarshal([]byte(request.Body), &ballSimulationRequest); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode:      400,
			IsBase64Encoded: false,
			Body:            "Invalid request: JSON must have `ball_count` and optionally `iteration_count`.",
		}, err
	}
	// Ensure the ballcount is in a valid range
	ballCount := ballSimulationRequest.BallCount
	if ballCount < 27 || ballCount > 127 {
		return events.APIGatewayProxyResponse{
			StatusCode:      400,
			IsBase64Encoded: false,
			Body:            "Invalid request: `ball_count` can't be less than 27 or greater than 127",
		}, InvalidBallCount
	}
	// Check to see if we're in mode-2
	if iterationCount := ballSimulationRequest.IterationCount; iterationCount != nil {
		ballClock := NewBallClock(ballCount)
		ballClock.TickMinutes(*iterationCount)

		response := BallClockSimulationResponse{
			BallClock: ballClock,
		}

		responseBytes, err := json.Marshal(response)
		if err != nil {
			return events.APIGatewayProxyResponse{
				StatusCode:      500,
				IsBase64Encoded: false,
				Body:            err.Error(),
			}, err
		}

		return events.APIGatewayProxyResponse{
			StatusCode:      200,
			IsBase64Encoded: false,
			Body:            string(responseBytes),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}, nil
	}

	// We're in mode-1
	clock, days, seconds := CalculateDaysUntilReset(ballCount)
	response := BallClockSimulationResponse{
		BallClock: *clock,
		Days:      &days,
		Seconds:   &seconds,
	}
	//
	responseBytes, err := json.Marshal(response)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode:      500,
			IsBase64Encoded: false,
			Body:            err.Error(),
		}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            string(responseBytes),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}

func main() {
	lambda.Start(Handler)
}
