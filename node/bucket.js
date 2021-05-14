class Bucket {

    constructor() {
        this.db = []
        for (let i = 0; i < 1e6; i++) {
            this.db.push({
                key: i,
                value: "value"
            })
        }
    }

    Get(key) {
        let len = this.db.length
        if (len === 0) {
            return false
        }
        let low = 0
        let high = len - 1
        let i = 0
        while (low <= high) {
            i++
            let mid = Math.floor((high + low) / 2)
            let item = this.db[mid]
            let val = item.key
            if (val === key)
                return i
            if (key < val)
                high = mid - 1
            if (key > val)
                low = mid + 1
        }
        return null
    }
}

module.exports = Bucket