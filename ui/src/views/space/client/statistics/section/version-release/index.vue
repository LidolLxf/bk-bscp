<template>
  <div>
    <SectionTitle :title="$t('配置版本发布')" />
    <Teleport :disabled="!isOpenFullScreen" to="body">
      <div
        :class="{ fullscreen: isOpenFullScreen }"
        @mouseenter="isShowOperationBtn = true"
        @mouseleave="isShowOperationBtn = false">
        <Card :title="$t('目标配置版本')" :height="344">
          <template #operation>
            <OperationBtn
              v-show="isShowOperationBtn"
              :is-open-full-screen="isOpenFullScreen"
              @refresh="loadChartData"
              @toggle-full-screen="isOpenFullScreen = !isOpenFullScreen" />
          </template>
          <template #head-suffix>
            <TriggerBtn v-model:current-type="currentType" style="margin-left: 16px" />
          </template>
          <bk-loading class="loading-wrap" :loading="loading">
            <component
              v-if="data?.length"
              :bk-biz-id="bkBizId"
              :app-id="appId"
              :is="currentComponent"
              :data="data"
              :is-full-screen="isOpenFullScreen"
              @update="jumpVersionName = $event as string"
              @jump="jumpToSearch" />
            <bk-exception v-else type="empty" scene="part" :description="$t('暂无数据')">
              <template #type>
                <span class="bk-bscp-icon icon-pie-chart exception-icon" />
              </template>
            </bk-exception>
          </bk-loading>
        </Card>
      </div>
    </Teleport>
  </div>
</template>

<script lang="ts" setup>
  import { ref, onMounted, computed, watch } from 'vue';
  import SectionTitle from '../../components/section-title.vue';
  import Card from '../../components/card.vue';
  import TriggerBtn from '../../components/trigger-btn.vue';
  import Pie from './pie.vue';
  import Column from './column.vue';
  import Table from './table.vue';
  import OperationBtn from '../../components/operation-btn.vue';
  import { IClientConfigVersionItem, IClinetCommonQuery } from '../../../../../../../types/client';
  import { getConfigVersionData } from '../../../../../../api/client';
  import useClientStore from '../../../../../../store/client';
  import { storeToRefs } from 'pinia';
  import { useRouter } from 'vue-router';

  const router = useRouter();

  const clientStore = useClientStore();
  const { searchQuery } = storeToRefs(clientStore);

  const props = defineProps<{
    bkBizId: string;
    appId: number;
  }>();

  const currentType = ref('pie');
  const componentMap = {
    pie: Pie,
    column: Column,
    table: Table,
  };
  const data = ref<IClientConfigVersionItem[]>();
  const currentComponent = computed(() => componentMap[currentType.value as keyof typeof componentMap]);
  const loading = ref(false);
  const jumpVersionName = ref('');
  const isOpenFullScreen = ref(false);
  const isShowOperationBtn = ref(false);

  onMounted(() => {
    loadChartData();
  });

  watch(
    () => props.appId,
    () => {
      loadChartData();
    },
  );

  watch(
    () => searchQuery.value,
    () => {
      loadChartData();
    },
    { deep: true },
  );

  const loadChartData = async () => {
    const params: IClinetCommonQuery = {
      last_heartbeat_time: searchQuery.value.last_heartbeat_time,
      search: searchQuery.value.search,
    };
    try {
      loading.value = true;
      const res = await getConfigVersionData(props.bkBizId, props.appId, params);
      data.value = res.client_config_version;
    } catch (error) {
      console.error(error);
    } finally {
      loading.value = false;
    }
  };

  const jumpToSearch = () => {
    const routeData = router.resolve({
      name: 'client-search',
      params: { appId: props.appId, bizId: props.bkBizId },
      query: { target_release_name: jumpVersionName.value, heartTime: searchQuery.value.last_heartbeat_time },
    });
    window.open(routeData.href, '_blank');
  };
</script>

<style scoped lang="scss">
  .fullscreen {
    position: fixed;
    top: 0;
    left: 0;
    width: 100vw;
    height: 100vh;
    z-index: 5000;
    .card {
      height: 100vh !important;
      :deep(.operation-btn) {
        top: 0 !important;
      }
    }
  }
  .loading-wrap {
    height: 100%;
  }

  :deep(.bk-exception) {
    height: 100%;
    justify-content: center;
    transform: translateY(-20px);
  }
</style>
