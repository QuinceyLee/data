package net.sppan.base.service.impl;

import net.sppan.base.dao.IAttackMsgDao;
import net.sppan.base.dao.support.IBaseDao;
import net.sppan.base.entity.AttackedMsg;
import net.sppan.base.service.IAttackedMsgService;
import net.sppan.base.service.support.impl.BaseServiceImpl;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class IAttackedMsgServiceImpl extends BaseServiceImpl<AttackedMsg,Integer>implements IAttackedMsgService {

   @Autowired
   private IAttackMsgDao iAttackMsgDao;

    @Override
    public IBaseDao<AttackedMsg, Integer> getBaseDao() {
        return iAttackMsgDao;
    }
}
