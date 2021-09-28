package handlers

// nolint

//func (c *controller) webhook(w http.ResponseWriter, r *http.Request) {
//	if r.Method != http.MethodPost {
//		http.Error(w, "POST requests only allowed", http.StatusMethodNotAllowed)
//		return
//	}
//
//	var objmap map[string]interface{}
//	err := json.NewDecoder(r.Body).Decode(&objmap)
//	defer closeBody(r.Body)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//	log.Println(objmap)
//}
//
//func (c *controller) newBot(w http.ResponseWriter, r *http.Request) {
//	if r.Method != http.MethodPost {
//		http.Error(w, "POST requests only allowed", http.StatusMethodNotAllowed)
//		return
//	}
//
//	token := c.auth.GetServiceTokenFromContext(r.Context())
//	if token == nil {
//		http.Error(w, "unauthorized", http.StatusUnauthorized)
//		return
//	}
//
//	var data struct {
//		Token string `json:"token"`
//		Type  string `json:"type"`
//	}
//
//	decoder := json.NewDecoder(r.Body)
//	defer closeBody(r.Body)
//
//	if err := decoder.Decode(&data); err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	bot := models.NewBot(data.Token, data.Type, token.GetTenantID())
//	if err := c.connectBot(bot); err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	if err := c.repo.UpsertBot(bot); err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	if err := prepareResponse(w, bot); err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	w.WriteHeader(http.StatusCreated)
//}
//
//func (c controller) send(w http.ResponseWriter, r *http.Request) {
//	if r.Method != http.MethodPost {
//		http.Error(w, "POST requests only allowed", http.StatusMethodNotAllowed)
//		return
//	}
//
//	decoder := json.NewDecoder(r.Body)
//	defer closeBody(r.Body)
//
//	var req struct {
//		Receiver string `json:"receiver"`
//		Text     string `json:"text"`
//	}
//
//	err := decoder.Decode(&req)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//
//	msg := models.NewMessage("", req.Text)
//	bot, err := c.repo.GetBotByID(mux.Vars(r)["bot"])
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusNotFound)
//		return
//	}
//	err = c.sendMessage(bot, msg, req.Receiver)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//	w.WriteHeader(http.StatusOK)
//}
//
//func (c controller) sendMessage(bot *models.Bot, msg *models.Message, receiver string) error {
//	if bot.IsTelegramBot() {
//		return c.tg.SendMessage(msg, bot.Token, receiver)
//	}
//	if bot.IsViberBot() {
//		return c.vb.SendMessage(msg, bot.Token, receiver)
//	}
//	return fmt.Errorf("unsupported platform")
//}

//func (c *controller) ConnectBot(bot *models.Bot) error {
//	if bot.IsTelegramBot() {
//		path := c.domain + s.getEndpointForTgBot(bot.ID)
//		return s.tg.ConnectNewBot(bot.Token, path)
//	}
//	if bot.IsViberBot() {
//		path := s.domain + s.getEndpointForVbBot(bot.ID)
//		return s.vb.ConnectNewBot(bot.Token, path)
//	}
//	return fmt.Errorf("unsupported platform")
//}
//
//func (c *controller) getEndpointForTgBot(id string) string {
//	return telegramEndpoint + id
//}
//
//func (c *controller) getEndpointForVbBot(id string) string {
//	return viberEndpoint + id
//}
