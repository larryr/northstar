local output = require("nsOutput")
local ftp = require("nsFTP")

function main()
    local connection, err = ftp.connect("10.37.2.6:21")
    if err ~= nil then
        error(err)
    end

    err = connection:login("test", "test")
    if err ~= nil then
        error(err)
    end

    err = connection:mkdir("/test")
    if err ~= nil then
        error(err)
    end

    err = connection:logout()
    if err ~= nil then
        error(err)
    end

    err = connection:disconnect()
    if err ~= nil then
        error(err)
    end
end