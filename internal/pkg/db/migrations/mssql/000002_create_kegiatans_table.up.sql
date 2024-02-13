IF NOT EXISTS
        (SELECT *FROM sysobjects
            WHERE id = object_id(N'dbo.kegiatans')
            AND OBJECTPROPERTY(id, N'IsUserTable') = 1
        )
        CREATE TABLE dbo.kegiatans (
            ID      int IDENTITY(1,1) PRIMARY KEY,
            Title   varchar(255) NOT NULL,
            Status  varchar(255) NOT NULL,
            UserID  int FOREIGN KEY REFERENCES Users(ID),
        );