package net.sppan.base.controller.admin.monitor;


import net.sppan.base.controller.BaseController;
import net.sppan.base.dao.IMachineInfoDao;
import net.sppan.base.entity.AttackedMsg;
import net.sppan.base.entity.MachineInfo;
import net.sppan.base.service.IAttackedMsgService;
import net.sppan.base.service.IMachineInfoService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.domain.Page;
import org.springframework.stereotype.Controller;
import org.springframework.ui.ModelMap;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.servlet.ModelAndView;

import java.util.List;

@Controller
@RequestMapping("/admin/machineInfo")
public class MachineInfoController extends BaseController {

    @Autowired
    private IMachineInfoService iMachineInfoService;

    @RequestMapping("/list")
    @ResponseBody
    public Page<MachineInfo> list() {
//		SimpleSpecificationBuilder<Resource> builder = new SimpleSpecificationBuilder<Resource>();
//		String searchText = request.getParameter("searchText");
//		if(StringUtils.isNotBlank(searchText)){
//			builder.add("name", Operator.likeAll.name(), searchText);
//		}
        Page<MachineInfo> page = iMachineInfoService.findAll(getPageRequest());
        return page;
    }

    @RequestMapping("/list/findById")
    @ResponseBody
    public MachineInfo listById(@RequestParam(name = "id") Integer id){
//		SimpleSpecificationBuilder<Resource> builder = new SimpleSpecificationBuilder<Resource>();
//		String searchText = request.getParameter("searchText");
//		if(StringUtils.isNotBlank(searchText)){
//			builder.add("name", Operator.likeAll.name(), searchText);
//		}
        MachineInfo page = iMachineInfoService.find(id);
        return page;
    }

    @RequestMapping(value = "/add", method = RequestMethod.GET)
    public String add(ModelMap map) {
        return "admin/sc/form";
    }
    @RequestMapping(value = { "/index" })
    public String index() {
        return "admin/sc/index";
    }


}
