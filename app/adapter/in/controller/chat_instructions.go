package controller

import (
	nats_publihser "github.com/Ignaciojeria/einar/app/adapter/out/nats_publisher"
	"github.com/Ignaciojeria/einar/app/domain"
	"github.com/Ignaciojeria/einar/app/shared/archetype/container"
	einar "github.com/Ignaciojeria/einar/app/shared/archetype/echo_server"
	nats "github.com/Ignaciojeria/einar/app/shared/archetype/nats"

	"net/http"

	"github.com/labstack/echo/v4"
)

func init() {
	container.InjectInboundAdapter(func() error {
		einar.Echo.POST("/api/chat/instructions", chatInstructions)
		return nil
	})
}

func chatInstructions(c echo.Context) error {
	var instruction domain.Instruction
	if err := c.Bind(&instruction); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	instruction.BrokerTopic = "public-server"
	instruction.CustomerTopic = nats.EinarTopic
	if err := nats_publihser.PublishInstruction(c.Request().Context(), instruction); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "instruction sended")
}
