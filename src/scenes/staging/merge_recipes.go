package staging

// Merge usage:
//
// yellow worker +++
// yellow militia ++
// red worker +++
// red militia +++
// green worker ++++
// green militia ++
// blue worker +++
// blue militia ++
//
// Used:
// freighter: yellow worker + green worker
// redminer: yellow worker + red worker
// crippler: yellow militia + green militia
// fighter: red militia + green militia
// servo: yellow worker + blue worker
// repeller: red worker + blue worker
// recharger: green worker + blue worker
// repair: red worker + blue militia
// generator: green worker + yellow militia
// mortar: green worker + red militia
// antiair: red militia + blue militia
//
// Unused:
// yellow worker + red militia
// yellow worker + green militia
// yellow worker + blue militia
// red worker + green worker
// red worker + green militia
// red worker + yellow militia
// green worker + blue militia
// blue worker + red militia
// blue worker + green militia
// blue worker + yellow militia
// yellow militia + blue militia
// yellow militia + red militia
// green militia + blue militia
var tier2agentMergeRecipeList = []agentMergeRecipe{
	{
		agent1kind:    agentWorker,
		agent1faction: greenFactionTag,
		agent2kind:    agentMilitia,
		agent2faction: redFactionTag,
		result:        mortarAgentStats,
	},
	{
		agent1kind:    agentMilitia,
		agent1faction: redFactionTag,
		agent2kind:    agentMilitia,
		agent2faction: blueFactionTag,
		result:        mortarAgentStats,
	},
	{
		agent1kind:    agentWorker,
		agent1faction: blueFactionTag,
		agent2kind:    agentWorker,
		agent2faction: blueFactionTag,
		result:        rechargeAgentStats,
	},
	{
		agent1kind:    agentWorker,
		agent1faction: yellowFactionTag,
		agent2kind:    agentWorker,
		agent2faction: yellowFactionTag,
		result:        freighterAgentStats,
	},
	{
		agent1kind:    agentWorker,
		agent1faction: redFactionTag,
		agent2kind:    agentWorker,
		agent2faction: redFactionTag,
		result:        redminerAgentStats,
	},

	{
		agent1kind:    agentMilitia,
		agent1faction: redFactionTag,
		agent2kind:    agentMilitia,
		agent2faction: redFactionTag,
		result:        fighterAgentStats,
	},
	{
		agent1kind:    agentWorker,
		agent1faction: yellowFactionTag,
		agent2kind:    agentWorker,
		agent2faction: greenFactionTag,
		result:        servoAgentStats,
	},
	{
		agent1kind:    agentMilitia,
		agent1faction: greenFactionTag,
		agent2kind:    agentMilitia,
		agent2faction: greenFactionTag,
		result:        cripplerAgentStats,
	},
	{
		agent1kind:    agentWorker,
		agent1faction: blueFactionTag,
		agent2kind:    agentMilitia,
		agent2faction: blueFactionTag,
		result:        repellerAgentStats,
	},
	{
		agent1kind:    agentWorker,
		agent1faction: blueFactionTag,
		agent2kind:    agentMilitia,
		agent2faction: yellowFactionTag,
		result:        generatorAgentStats,
	},
	{
		agent1kind:    agentWorker,
		agent1faction: redFactionTag,
		agent2kind:    agentMilitia,
		agent2faction: greenFactionTag,
		result:        repairAgentStats,
	},
}

var tier3agentMergeRecipeList = []agentMergeRecipe{
	{
		agent1kind: agentRepeller,
		agent2kind: agentFreighter,
		evoCost:    5,
		result:     flamerAgentStats,
	},
	{
		agent1kind: agentFighter,
		agent2kind: agentFighter,
		evoCost:    11,
		result:     destroyerAgentStats,
	},

	{
		agent1kind: agentRecharger,
		agent2kind: agentRepair,
		evoCost:    7,
		result:     refresherAgentStats,
	},
}