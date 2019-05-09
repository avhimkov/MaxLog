// 		var global = {
// 			lang: "rus", // Язык интерфейса
// 			langPack: {}, // Тут хранится языковой пакет интерфейса

// 			languages: [], // Список языков с сервера
// 			language: 0, // ID выбранного языка

// 			workplaces: [], // Рабочие места
// 			workusers: [], // Список сотрудников
// 			services: [], // Услуги

// 			selectedTicket: null, // Выбранный талончик (для функции "перепечатать")

// 			lostConnection: false, // Подключение к серверу изначально есть (если его нет, то страница не откроется)

// 			updateTime: 5, // Интервал проверки сеодинения (в секундах)

// 			options: {
// 				custDataLength: 3,
// 				multiReg: false,
// 				barcode: false, // Печатать штрих-код (по умолчанию - не печатать)
// 				allowWU: false, // Использовать роли (по умолчанию - рабочие места)
// 				ticketState: [0, 5, 6], // Состояния посетителей, при запросе списка талонов
// 				priority: ["", "VIP"], // Приоретите
// 				usePIN: false, // Используется ли пин-код (по умолчанию - нет)
// 				registerSRToBusy: false, // разрешить регистрацию на занятый интервал
// 				sRPINLength: 9, //длина пинкода
// 				qrCode: "", // печать QR кода
// 				appVer: "", // версия сервера
// 				maxWaitingTime: 10 // максимальное время ожидания
// 			},

//  /**
//      * Асинхронная регистрация талона
//      * @param {{serviceId, languageId, custData, noteText, priorityId, userId}} params
//      * @returns promise
//      */
// multiRegisterAsync: function (params) {
// var dfr = $.Deferred();

// var data = {
// 	command: "cmd_Registers",
// 	ServiceCount: params.serviceId.length,
// 	LangID: params.languageId,
// 	CustData: params.custData,
// 	Note: params.noteText
// }

// 		multiRegister: function (callback) {
// 			"use strict";
// 			var data = {
// 				command: "cmd_Registers",
// 				ServiceCount: global.register.id.length,
// 				LangID: global.language,
// 				CustData: global.register.custData
// 			}

// 			for (var i = 0; i < global.register.id.length; i++) {
// 				data["ServiceID_" + (i + 1)] = global.register.id[i];
// 				data["PriorityID_" + (i + 1)] = global.register.priority;
// 				if (global.register.to[i] != -1) {
// 					data["RegToID_" + (i + 1)] = global.register.to[i];
// 				}

// 				if (data["ServiceID_" + (i + 1)] === undefined) data["ServiceID_" + (i + 1)] = '';
// 				if (data["PriorityID_" + (i + 1)] === undefined) data["PriorityID_" + (i + 1)] = '';
// 				if (data["RegToID_" + (i + 1)] === undefined) data["RegToID_" + (i + 1)] = '';
// 			}

// for (var i = 0; i < params.serviceId.length; i++) {
// 	data["ServiceID_" + (i + 1)] = params.serviceId[i];
// 	data["PriorityID_" + (i + 1)] = params.priorityId;
// 	if (global.register.to[i] != -1) {
// 		data["RegToID_" + (i + 1)] = params.userId[i];
// 	}

// 	if (data["ServiceID_" + (i + 1)] === undefined) data["ServiceID_" + (i + 1)] = '';
// 	if (data["PriorityID_" + (i + 1)] === undefined) data["PriorityID_" + (i + 1)] = '';
// 	if (data["RegToID_" + (i + 1)] === undefined) data["RegToID_" + (i + 1)] = '';
// }