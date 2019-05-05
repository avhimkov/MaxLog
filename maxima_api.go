package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

//http://" + location.hostname + ":5612/QMSSR

func maxNewTicket(ip string) {
	// Set variable values
	gatewayURL := "http://" + ip + ":5612/QMSSR"
	urlStr := gatewayURL

	// Params
	v := url.Values{}
	v.Set("ip", ip)
	rb := *strings.NewReader(v.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &rb)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make request
	resp, _ := client.Do(req)

	//print response
	fmt.Println(resp.Status)

}


// Получение списка языков
"cmd_GetLanguages"

global.languages[i] = {
	id: res.Languages[i].ID,
	name: res.Languages[i].Name,
	short: res.Languages[i].ShortName
};

 // получение параметров терминала (server.connect.config)
"command=cmd_getconfig"
global.options.barcode = res.PrintBarCode;
global.options.allowWU = res.UseWorkUsersRoles;

global.options.maxDays = res.SRDaysAfter; // кол-во дней для самозаписи
global.options.sRCurDay = res.SRCurDay; //разрешить запись на сегодня
global.options.sRPINLength = res.SRPINLength;

global.options.usePIN = res.Terminal_UsePIN;
global.options.qrCode = res.QRCodeText || "";
global.options.appVer = res.AppVer || "";

global.options.maxWaitingTime = res.CriticalWaitingTime || 10;

// Получение рабочих мест
"cmd_GetWorkplaces"
global.workplaces[i] = {
	id: res.Workplaces[i].ID,
	idCon: res.Workplaces[i].IDCon,
	name: res.Workplaces[i].Name,
	placeNo: res.Workplaces[i].PlaceNo || "",
	placeNo2: res.Workplaces[i].PlaceNo2 || ""
};

 // Получение рабочих мест
"cmd_GetWorkusers"
global.workusers[i] = {
	id: res.WorkUsers[i].ID,
	name: res.WorkUsers[i].FirstName || "",
	last: res.WorkUsers[i].LastName || "",
	second: res.WorkUsers[i].Patronimyc || ""
};

/**
* Асинхронный запрос услуг
* @returns promise
*/
	 getServicesAsync: function () {
        "use strict";
        var dfr = $.Deferred();
        var data = {
            command: "cmd_getservices",
            TypeOfMenu: 2,
            LanguageID: global.language
		}
		var result = {
			code: res.ResultCode,
			command: $.param(data)
		};

/**
* Асинхронная регистрация талона
* @param {{serviceId, languageId, custData, noteText, priorityId, userId}} params
* @returns promise
*/
command: "cmd_Register",
ServiceID: params.serviceId,
LangID: params.languageId,
CustData: params.custData,
Note: params.noteText,
PriorityID: params.priorityId,
RegToID: params.userId == -1 ? undefined : params.userId
if (params.serviceId instanceof Array) {
	data = {
		command: 'cmd_Registers',
		ServiceCount: params.serviceId.length,
		LangID: params.languageId,
		CustData: params.custData,
	};

	for (i = 1; i <= params.serviceId.length; i++) {
		data['ServiceID_' + i] = params.serviceId[i - 1];
		data['RegToID_' + i] = params.userId[i - 1];
	}
}


/** Асинхронно получить информацию по пину
     * [registerByPIN description]
     * @param {number}   pin      [description]
     * @param {[type]}   attr     [description]
     * @param {Function} callback [description]
     */
	 шnformationByPIN: function (params) {
        var dfr = $.Deferred();

        var data = {
            command: 'cmd_SR_InformationByPIN',
            LangID: params.languageId,
            PIN: params.pin
		};
		
	/** Асинхронная регистрация по пину
* [registerByPIN description]
* @param {number}   pin      [description]
* @param {[type]}   attr     [description]
* @param {Function} callback [description]
*/
registerByPIN: function (params) {
var dfr = $.Deferred();

var data = {
	command: 'cmd_SR_RegisterByPIN',
	LangID: params.languageId,
	PIN: params.pin
};


 /**
     * Асинхронная регистрация талона
     * @param {{serviceId, languageId, custData, noteText, priorityId, userId}} params
     * @returns promise
     */
multiRegisterAsync: function (params) {
var dfr = $.Deferred();

var data = {
	command: "cmd_Registers",
	ServiceCount: params.serviceId.length,
	LangID: params.languageId,
	CustData: params.custData,
	Note: params.noteText
}


for (var i = 0; i < params.serviceId.length; i++) {
	data["ServiceID_" + (i + 1)] = params.serviceId[i];
	data["PriorityID_" + (i + 1)] = params.priorityId;
	if (global.register.to[i] != -1) {
		data["RegToID_" + (i + 1)] = params.userId[i];
	}

	if (data["ServiceID_" + (i + 1)] === undefined) data["ServiceID_" + (i + 1)] = '';
	if (data["PriorityID_" + (i + 1)] === undefined) data["PriorityID_" + (i + 1)] = '';
	if (data["RegToID_" + (i + 1)] === undefined) data["RegToID_" + (i + 1)] = '';
}

 // Получение информации об услуге
    /**
     * @param int id
     * @returns promise
     */
	getServiceByIdAsync: function (id) {
	var dfr = $.Deferred();
	var data = {
		command: 'cmd_GetServiceByID',
		GetHandlingWPorWU: 1,
		ServiceID: id,
		GetWaitingCountPerWPOrWU: 1
	}

	if (res.ResultCode === "0") {
		var waitingList = res.WaitingCountWP || res.WaitingCountWU;
		var handling = [];
		if (res.HandlingWP) {
			global.handlingWP = true;
			handling = res.HandlingWP.sort(function (a, b) {
				return a.Name > b.Name;
			});
		} else if (res.HandlingWU) {
			global.handlingWP = false;
			handling = res.HandlingWU.sort(function (a, b) {
				return [a.LastName, a.FirstName, a.Patronymic].join() > [b.LastName, b.FirstName, b.Patronymic].join();
			});
		}

		dfr.resolve(handling, waitingList, res);

		 /**
     * Асинхронная регистрация талона на время
     * @param {{serviceId, languageId, custData, noteText, priorityId, userId}} params
     * @returns promise
     */
	 registerSRAsync: function (params) {
        "use strict";
        var dfr = $.Deferred();
        var data = {
            command: "cmd_SR_Register",
            ServiceID: params.serviceId,
            LangID: params.languageId,
            CustData: params.custData,
            Note: params.note,
            PriorityID: params.priorityId,
            CallTime: params.callTime,
            RegToID: params.userId == -1 ? undefined : params.userId
		}
		
		 /**
     * Асинхронная регистрация талона на время
     * @param {{serviceId, languageId, custData, noteText, priorityId, userId}} params
     * @returns promise
     */
	 getIntervalsByServiceIdAsync: function (params) {

        var dfr = $.Deferred();
        var data = {
            command: "cmd_GetIntervals",
            ServiceID: params.id,
            DateFrom: params.date,
            DateTo: params.dateTo || utils.incdate(params.date, 1),
            RegToID: params.userId == -1 ? undefined : params.userId
		}
		
		/**
     * Асинхронное удаление талона в средварительной регистрации
     * @param {number} id
     * @returns promise
     */
	 removeTicketSr: function (id) {
        var dfr = $.Deferred();
        var data = {
            command: "cmd_SRDelTicketStep",
            SRTicketStepID: id,
		}
		
	// Получение списка услуг
	getServices: function () {
	"use strict";
	var data = {
		command: "cmd_getservices",
		TypeOfMenu: 2,
		LanguageID: global.language
	};
	updateTicketStep: function (options, callback) {
        var data = {
            command: 'cmd_UpdateTicketStep',
            TicketStepID: undefined,
            ServiceID: undefined,
            PriorityID: undefined,
            RatingID: undefined,
            LangID: undefined,
            CustData: undefined,
            Note: undefined,
            RegToID: undefined,
            State: undefined
		};
		
	updateSRTicketStep: function (options, callback) {

		var data = {
			command: "cmd_SRUpdateTicketStep",
			SRTicketStepID: undefined,
			// ServiceID: undefined,
			// CustData: undefined,
			// Note: undefined,
			// CallTime: undefined
		};

		// Получение инфорамии об услуге
		getServiceById: function (id, callback) {
		"use strict";

		var data = {
			command: 'cmd_GetServiceByID',
			GetHandlingWPorWU: 1,
			ServiceID: id,
			GetWaitingCountPerWPOrWU: 1
		}

		if (res.ResultCode === "0") {
			var waitingList = res.WaitingCountWP || res.WaitingCountWU;
			var handling = res.HandlingWP || res.HandlingWP;
			callback(handling, waitingList);
		} else {
			server.log({
				code: res.ResultCode,
				command: $.param(data)
			});
		}

		error: function () {
			server.log({
				code: 820
			});
		}

		/**
     * Регистрация талона
     */
	 register: function (callback) {
        "use strict";
        var cmd = "command=cmd_Register&ServiceID=" +
            global.register.id +
            "&LangID=" + global.language +
            "&CustData=" + global.register.custData +
            "&Note=" + global.register.note +
            "&PriorityID=" + global.register.priority;

        if (global.register.to !== -1) {
            cmd += "&RegToID=" + global.register.to;
		}
		

		multiRegister: function (callback) {
			"use strict";
			var data = {
				command: "cmd_Registers",
				ServiceCount: global.register.id.length,
				LangID: global.language,
				CustData: global.register.custData
			}
	
			for (var i = 0; i < global.register.id.length; i++) {
				data["ServiceID_" + (i + 1)] = global.register.id[i];
				data["PriorityID_" + (i + 1)] = global.register.priority;
				if (global.register.to[i] != -1) {
					data["RegToID_" + (i + 1)] = global.register.to[i];
				}
	
				if (data["ServiceID_" + (i + 1)] === undefined) data["ServiceID_" + (i + 1)] = '';
				if (data["PriorityID_" + (i + 1)] === undefined) data["PriorityID_" + (i + 1)] = '';
				if (data["RegToID_" + (i + 1)] === undefined) data["RegToID_" + (i + 1)] = '';
			}

		// Регистрация талона на предварительную запись
		registerSR: function (callback) {
			"use strict";
			var data = {
				command: "cmd_SR_Register",
				ServiceID: global.register.id,
				LangID: global.language,
				CustData: global.register.custData,
				Note: global.register.note,
				CallTime: global.register.callTime,
				SR_PIN: 1
			}
	
			if (global.register.to !== -1) {
				data.RegToID = global.register.to;
			}

			/** Соединение с сервером востановлено */
			restoreConnection: function () {
				if (global.lostConnection) {
					global.lostConnection = false;
					actions.connection.online();
					uiConstructor.make.overlay.hide()
				}
			},
		
			updateCheckConnection: function () {
				if (this.connectionTimeout) {
					clearTimeout(this.connectionTimeout);
				}
		
				this.connectionTimeout = setTimeout(function () {
					server.checkConnection();
					this.connectionTimeout = null;
				}, global.updateTime * 1000);
			},
		
			// Проверка соединения с сервером
			checkConnection: function () {
				var self = this;
				$.ajax({
					async: true,
					dataType: "json",
				});
			},

			// Получение списка талонов
			getTickets: function (state, callback) {
			"use strict";
	
			var data = {
				command: "cmd_GetTicketSteps",
				State: state.join(",")
			}
			if (state.length === 0) {
				callback([]);

			// Получение списка талонов по предварительной записи
			getTicketsSR: function (startDate, endDate, callback) {
				"use strict";
				var data = {
					command: "cmd_GetSRTicketSteps",
					SRDate: startDate
				}
		
				if (endDate) {
					data["SRDateTo"] = endDate
				}

			getIntervalsByServiceId: function (date) {
				"use strict";
				var data = {
					command: "cmd_GetIntervals",
					ServiceID: global.register.id,
					DateFrom: date,
					// ToDo: если в месяце 30 дней
					//DateTo: date//utils.incdate(date, 1)
				}

		/*
		multiRegister() - если в первой выбранной услуге указываться RegToId, то у всех последующих
		выбранных услугах, у которых RegToId не указан будет тот же RegToId, что и у первой услуги.

		*/

		var global = {
			lang: "rus", // Язык интерфейса
			langPack: {}, // Тут хранится языковой пакет интерфейса
		
			languages: [], // Список языков с сервера
			language: 0, // ID выбранного языка
		
			workplaces: [], // Рабочие места
			workusers: [], // Список сотрудников
			services: [], // Услуги
		
			selectedTicket: null, // Выбранный талончик (для функции "перепечатать")
		
			lostConnection: false, // Подключение к серверу изначально есть (если его нет, то страница не откроется)
		
			updateTime: 5, // Интервал проверки сеодинения (в секундах)
		
			options: {
				custDataLength: 3,
				multiReg: false,
				barcode: false, // Печатать штрих-код (по умолчанию - не печатать)
				allowWU: false, // Использовать роли (по умолчанию - рабочие места)
				ticketState: [0, 5, 6], // Состояния посетителей, при запросе списка талонов
				priority: ["", "VIP"], // Приоретите
				usePIN: false, // Используется ли пин-код (по умолчанию - нет)
				registerSRToBusy: false, // разрешить регистрацию на занятый интервал
				sRPINLength: 9, //длина пинкода
				qrCode: "", // печать QR кода
				appVer: "", // версия сервера
				maxWaitingTime: 10 // максимальное время ожидания
			},
		
			register: {
				id: null, // ID-услуги, на которую будет производиться регистрация
				custData: "", // Информация о посетителе
				note: "", // Примечание оператора
				priority: 0, // Приоритет посетителя
				to: -1, // ID рабочего места или роли, к кому направить посетителя (по умолчанию(-1) - к любому)
				callTime: ''
			},