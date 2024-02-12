IF NOT EXISTS
        (SELECT *FROM sysobjects
            WHERE id = object_id(N'dbo.users')
            AND OBJECTPROPERTY(id, N'IsUserTable') = 1
        )
        CREATE TABLE dbo.users (
            ID         int IDENTITY(1,1) PRIMARY KEY,
            Username   varchar(255) NOT NULL,
            Password   varchar(255) NOT NULL
        );