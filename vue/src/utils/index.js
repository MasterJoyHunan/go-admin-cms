/**
 * 删除空的
 */
export function removeEmptyChildren(data, key = 'children') {
    for (const v of data) {
        if (v[key].length > 0) {
            removeEmptyChildren(v[key], key)
        } else {
            delete v[key]
        }
    }
    return data
}
