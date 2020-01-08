package model

// Symbol represents a information fundamental from corporation
type Symbol struct {
	Papel                string
	Cotacao              string
	Tipo                 string
	DataUtlCot           string
	Empresa              string
	Setor                string
	Subsetor             string
	Min52Sem             string
	Max52Sem             string
	VolMed               string
	ValorMercado         string
	UltBalancoProcessado string
	ValorDaFirma         string
	NroAcoes             string
	PL                   string
	LPA                  string
	PVP                  string
	VPA                  string
	PEBIT                string
	MargBruta            string
	Dozemeses            string
	PSR                  string
	MargEBIT             string
	PAtivos              string
	MargLiquida          string
	ReceitaLiquida       string
	PCapGiro             string
	EBIT                 string
	PAtivCircLiq         string
	ROIC                 string
	DivYield             string
	ROE                  string
	EVEBITDA             string
	LiquidezCorrente     string
	EVEBIT               string
	DivBrPatrim          string
	CresRec5a            string
	GiroAtivos           string
	Ativo                string
	Depositos            string
	DividaLiquida        string
	CartdeCredito        string
	PatrimLiq            string
	ResultIntFinanc      string
	RecServicos          string
	LucroLiquido         string
	AtivoCirculante      string
	DividaBruta          string
	Disponibilidades     string
}

// Build a symbol entity
func Build(symbol *Symbol, k, v string) {
	switch k {
	case "Papel":
		symbol.Papel = v
	case "Vol $ méd (2m)":
		symbol.VolMed = v
	case "P/Cap. Giro":
		symbol.PCapGiro = v
	case "EV / EBITDA":
		symbol.EVEBITDA = v
	case "Receita Líquida":
		symbol.ReceitaLiquida = v
	case "Valor da firma":
		symbol.ValorDaFirma = v
	case "P/VP":
		symbol.PVP = v
	case "Dív. Líquida":
		symbol.DividaLiquida = v
	case "Lucro Líquido":
		symbol.LucroLiquido = v
	case "Min 52 sem":
		symbol.Min52Sem = v
	case "Nro. Ações":
		symbol.NroAcoes = v
	case "P/Ativ Circ Liq":
		symbol.PAtivCircLiq = v
	case "Div. Yield":
		symbol.DivYield = v
	case "Cres. Rec (5a)":
		symbol.CresRec5a = v
	case "Ativo Circulante":
		symbol.AtivoCirculante = v
	case "EBIT":
		symbol.EBIT = v
	case "Empresa":
		symbol.Empresa = v
	case "P/L":
		symbol.PL = v
	case "P/Ativos":
		symbol.PAtivos = v
	case "EV / EBIT":
		symbol.EVEBIT = v
	case "Ativo":
		symbol.Ativo = v
	case "Dív. Bruta":
		symbol.DividaBruta = v
	case "Cotação":
		symbol.Cotacao = v
	case "Data últ cot":
		symbol.DataUtlCot = v
	case "P/EBIT":
		symbol.PEBIT = v
	case "Patrim. Líq":
		symbol.PatrimLiq = v
	case "Setor":
		symbol.Setor = v
	case "Valor de mercado":
		symbol.ValorMercado = v
	case "PSR":
		symbol.PSR = v
	case "Max 52 sem":
		symbol.Max52Sem = v
	case "Últ balanço processado":
		symbol.UltBalancoProcessado = v
	case "Disponibilidades":
		symbol.Disponibilidades = v
	case "Tipo":
		symbol.Tipo = v
	case "Subsetor":
		symbol.Subsetor = v
	case "ROE":
		symbol.ROE = v
	case "ROIC":
		symbol.ROIC = v
	case "Marg. EBIT":
		symbol.MargEBIT = v
	case "Marg. Líquida":
		symbol.MargLiquida = v
	case "Giro Ativos":
		symbol.PCapGiro = v
	case "VPA":
		symbol.VPA = v
	case "Div Br/ Patrim":
		symbol.DivBrPatrim = v
	case "Marg. Bruta":
		symbol.MargBruta = v
	}
}
