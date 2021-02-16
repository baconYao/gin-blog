# gin-blog

## Prerequisite (MySQL DB)

### Run a MySQL DB by using Docker
```bash
# user: root
# password: root
docker run --name gin-blog-service -e MYSQL_ROOT_PASSWORD=root -p 3306:3306 -d mysql:8.0

# enter mysql's interaction mode by (-u) root identity and (-p) without pointing a specific database
docker exec -it gin-blog-service mysql -u root -p
```

### Create a `blog_service` database
``` mysql
CREATE DATABASE
 IF
   NOT EXISTS blog_service DEFAULT CHARACTER
   SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;

mysql> show databases;
+--------------------+
| Database           |
+--------------------+
| blog_service       |
| information_schema |
| mysql              |
| performance_schema |
| sys                |
+--------------------+
5 rows in set (0.00 sec)
```

### Create a `blog_tag` table
```mysql
create table `blog_tag` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(100) DEFAULT '' COMMENT 'Tag name',
    `created_on` int(10) unsigned DEFAULT '0' COMMENT 'Create time',
    `created_by` varchar(100) DEFAULT '' COMMENT 'Owner',
    `modified_on` int(10) unsigned DEFAULT '0' COMMENT 'Modified time',
    `modified_by` varchar(100) DEFAULT '' COMMENT 'Modifier',
    `deleted_on` int(10) unsigned DEFAULT '0' COMMENT 'Delete time',
    `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '0 is not be deleted yet otherwise 1',
    `state` tinyint(3) unsigned DEFAULT '1' COMMENT '0 is disable, 1 is enable',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Tag management';
```

### Create a `blog_article` table
```mysql
create table `blog_article` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `title` varchar(100) DEFAULT '' COMMENT 'Article name',
    `desc` varchar(255) DEFAULT '' COMMENT 'Article description',
    `cover_image_url` varchar(255) DEFAULT '' COMMENT 'URL address of coverd image',
    `content` longtext COMMENT 'Article content',
    `created_on` int(10) unsigned DEFAULT '0' COMMENT 'Create time',
    `created_by` varchar(100) DEFAULT '' COMMENT 'Owner',
    `modified_on` int(10) unsigned DEFAULT '0' COMMENT 'Modified time',
    `modified_by` varchar(100) DEFAULT '' COMMENT 'Modifier',
    `deleted_on` int(10) unsigned DEFAULT '0' COMMENT 'Delete time',
    `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '0 is not be deleted yet otherwise 1',
    `state` tinyint(3) unsigned DEFAULT '1' COMMENT '0 is disable, 1 is enable',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Article management';
```

### Create a `blog_article_tag` table
```mysql
create table `blog_article_tag` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `article_id` int(11) NOT NULL COMMENT 'Article ID',
    `tag_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT 'Tag ID',
    `created_on` int(10) unsigned DEFAULT '0' COMMENT 'Create time',
    `created_by` varchar(100) DEFAULT '' COMMENT 'Owner',
    `modified_on` int(10) unsigned DEFAULT '0' COMMENT 'Modified time',
    `modified_by` varchar(100) DEFAULT '' COMMENT 'Modifier',
    `deleted_on` int(10) unsigned DEFAULT '0' COMMENT 'Delete time',
    `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '0 is not be deleted yet otherwise 1',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Link between article and tag';
```

## How to regenerate Swagger API Doc?

After `go get`, typing `swag init` at your terminal, you will get an updated *docs* folder
