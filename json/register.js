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

// /**
// * Асинхронная регистрация талона
// * @param {{serviceId, languageId, custData, noteText, priorityId, userId}} params
// * @returns promise
// */
// command: "cmd_Register",
// ServiceID: params.serviceId,
// LangID: params.languageId,
// CustData: params.custData,
// Note: params.noteText,
// PriorityID: params.priorityId,
// RegToID: params.userId == -1 ? undefined : params.userId
// if (params.serviceId instanceof Array) {
// 	data = {
// 		command: 'cmd_Registers',
// 		ServiceCount: params.serviceId.length,
// 		LangID: params.languageId,
// 		CustData: params.custData,
// 	};

// 	for (i = 1; i <= params.serviceId.length; i++) {
// 		data['ServiceID_' + i] = params.serviceId[i - 1];
// 		data['RegToID_' + i] = params.userId[i - 1];
// 	}
// }

// 		/**
//      * Регистрация талона
//      */
// 	 register: function (callback) {
//         "use strict";
//         var cmd = "command=cmd_Register&ServiceID=" +
//             global.register.id +
//             "&LangID=" + global.language +
//             "&CustData=" + global.register.custData +
//             "&Note=" + global.register.note +
//             "&PriorityID=" + global.register.priority;

//         if (global.register.to !== -1) {
//             cmd += "&RegToID=" + global.register.to;
// 		}

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