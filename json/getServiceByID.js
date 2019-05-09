// 		if (res.ResultCode === "0") {
// 			var waitingList = res.WaitingCountWP || res.WaitingCountWU;
// 			var handling = res.HandlingWP || res.HandlingWP;
// 			callback(handling, waitingList);
// 		} else {
// 			server.log({
// 				code: res.ResultCode,
// 				command: $.param(data)
// 			});
// 		}

// 	if (res.ResultCode === "0") {
// 		var waitingList = res.WaitingCountWP || res.WaitingCountWU;
// 		var handling = [];
// 		if (res.HandlingWP) {
// 			global.handlingWP = true;
// 			handling = res.HandlingWP.sort(function (a, b) {
// 				return a.Name > b.Name;
// 			});
// 		} else if (res.HandlingWU) {
// 			global.handlingWP = false;
// 			handling = res.HandlingWU.sort(function (a, b) {
// 				return [a.LastName, a.FirstName, a.Patronymic].join() > [b.LastName, b.FirstName, b.Patronymic].join();
// 			});
// 		}