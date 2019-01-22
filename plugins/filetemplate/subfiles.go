package pluginfiletemplate

import (
	"path"

	"github.com/zhs007/jarvistelebot/chatbotdb"
	"github.com/zhs007/jarvistelebot/chatbotdb/proto"
)

// procSubfiles - process subfiles
func procSubfiles(db *chatbotdb.ChatBotDB, uid string, jarvisnodename string, fd []byte, subfilesPath string) error {
	sf, err := chatbotdb.LoadSubfilesFromBuff(fd)
	if err != nil {
		return err
	}

	for _, v := range sf.Subfiles {
		us := &chatbotdbpb.UserFileTemplate{
			FileTemplateName: v,
			JarvisNodeName:   jarvisnodename,
			FullPath:         path.Join(subfilesPath, v),
		}

		err := db.SaveFileTemplate(uid, us)
		if err != nil {
			return err
		}
	}

	return nil
}
