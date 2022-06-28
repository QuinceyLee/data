package net.sppan.base.service.impl;

import net.sppan.base.dao.IAttackMsgDao;
import net.sppan.base.dao.IMachineInfoDao;
import net.sppan.base.dao.support.IBaseDao;
import net.sppan.base.entity.AttackedMsg;
import net.sppan.base.entity.MachineInfo;
import net.sppan.base.service.IAttackedMsgService;
import net.sppan.base.service.IMachineInfoService;
import net.sppan.base.service.support.impl.BaseServiceImpl;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;


@Service
public class IMachineInfoServiceImpl extends BaseServiceImpl<MachineInfo,Integer> implements IMachineInfoService {

    @Autowired
    private IMachineInfoDao iMachineInfoDao;

    @Override
    public IBaseDao<MachineInfo, Integer> getBaseDao() {
        return iMachineInfoDao;
    }
}




