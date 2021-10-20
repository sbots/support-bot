package repository

// nolint

// type BotsPlatform struct {
//}

// func (p *BotsPlatform) ConnectBot(bot *models.Bot) error {
//	if bot.IsTelegramBot() {
//		path := s.domain + s.getEndpointForTgBot(bot.ID)
//		return s.tg.ConnectNewBot(bot.Token, path)
//	}
//	if bot.IsViberBot() {
//		path := s.domain + s.getEndpointForVbBot(bot.ID)
//		return s.vb.ConnectNewBot(bot.Token, path)
//	}
//	return fmt.Errorf("unsupported platform")
//}
//
//func (p *BotsPlatform) getEndpointForTgBot(id string) string {
//	return telegramEndpoint + id
//}
//
//func (p *BotsPlatform) getEndpointForVbBot(id string) string {
//	return viberEndpoint + id
//}
