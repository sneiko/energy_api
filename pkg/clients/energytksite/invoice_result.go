package energytksite

type InvoiceResult struct {
	IdSending  int64  `json:"idSending"`
	Docnum     string `json:"docnum"`
	IdCityFrom int    `json:"idCityFrom"`
	CityFrom   struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"cityFrom"`
	IdCityTo int `json:"idCityTo"`
	CityTo   struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"cityTo"`
	WarehouseTo struct {
		Id          int    `json:"id"`
		Title       string `json:"title"`
		Address     string `json:"address"`
		Fulladdress string `json:"fulladdress"`
		WorkTime    []struct {
			Day   int    `json:"day"`
			Begin string `json:"begin"`
			End   string `json:"end"`
		} `json:"workTime"`
		Phone      string  `json:"phone"`
		Email      string  `json:"email"`
		Zipcode    int     `json:"zipcode"`
		Latitude   float64 `json:"latitude"`
		Longitude  float64 `json:"longitude"`
		IdCity     int     `json:"idCity"`
		IsInternal int     `json:"isInternal"`
		Code       string  `json:"code"`
		IsIssuer   int     `json:"isIssuer"`
		Type       int     `json:"type"`
		MaxWeight  int     `json:"maxWeight"`
		MaxVolume  int     `json:"maxVolume"`
		MaxLength  float64 `json:"maxLength"`
		MaxWidth   float64 `json:"maxWidth"`
		MaxHeight  float64 `json:"maxHeight"`
	} `json:"warehouseTo"`
	IdClientFrom  int64 `json:"idClientFrom"`
	IdClientTo    int64 `json:"idClientTo"`
	PriceTripType struct {
		Id    int    `json:"id"`
		Title string `json:"title"`
	} `json:"priceTripType"`
	Places                    int     `json:"places"`
	Weight                    int     `json:"weight"`
	Volume                    float64 `json:"volume"`
	DeliveryDateFrom          int     `json:"deliveryDateFrom"`
	DeliveryDateFromFormatted string  `json:"deliveryDateFromFormatted"`
	DeliveryDateTo            int     `json:"deliveryDateTo"`
	DeliveryDateToFormatted   string  `json:"deliveryDateToFormatted"`
	SenderTotalPrice          int     `json:"senderTotalPrice"`
	RecipientTotalPrice       int     `json:"recipientTotalPrice"`
	SenderIsPaid              bool    `json:"senderIsPaid"`
	RecipientIsPaid           bool    `json:"recipientIsPaid"`
	IsDelivery                bool    `json:"isDelivery"`
	States                    []struct {
		IdState             int    `json:"idState"`
		Title               string `json:"title"`
		IdSubState          int    `json:"idSubState"`
		SubStateTitle       string `json:"subStateTitle"`
		MovingDate          int    `json:"movingDate"`
		MovingDateFormatted string `json:"movingDateFormatted"`
		Amount              int    `json:"amount"`
		StateInfo           struct {
			IdWare    int         `json:"idWare"`
			Warehouse interface{} `json:"warehouse"`
			Trip      struct {
				Id       int64 `json:"id"`
				CityFrom struct {
					Id   int    `json:"id"`
					Name string `json:"name"`
				} `json:"cityFrom"`
				CityTo struct {
					Id   int    `json:"id"`
					Name string `json:"name"`
				} `json:"cityTo"`
				IdCityFrom int `json:"idCityFrom"`
				IdCityTo   int `json:"idCityTo"`
				IdWareFrom int `json:"idWareFrom"`
				IdWareTo   int `json:"idWareTo"`
				Date       int `json:"date"`
			} `json:"trip"`
			Issued interface{} `json:"issued"`
		} `json:"stateInfo"`
	} `json:"states"`
}
