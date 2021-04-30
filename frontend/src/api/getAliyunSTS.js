import request from '/@/utils/request'

let getMenus = {};

getMenus.send = function(data) {
    return request({
        url: `/api/aliyun/sts`,
        data,
        method: 'GET'
    })
}

export default getMenus