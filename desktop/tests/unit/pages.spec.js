import Pages from "@/components/Pages.vue";
import { mount } from "@vue/test-utils";
import { createStore } from "vuex";

describe("Pages.vue", () => {
  const createVuexStore = () => {
    return createStore({
      state() {
        return {
          Pages: {
            all: [{ ID: "1", title: "Post 1", body: "A simple page text body" }],
          },
        };
      },
      mutations: {
        updatePagesList(state) {
          state.Pages = [];
        },
      },
      actions: {
        listPages(ctx) {
          const res = mockGet();
          ctx.commit("updatePagesList", res);
        },
      },
    });
  };

  function factory() {
    const store = createVuexStore();
    return mount(Pages, {
      global: {
        plugins: [store],
      },
    });
  }

  const mockGet = jest.fn();

  jest.mock("axios", () => ({
    get: () => mockGet(),
  }));

  it("renders props when passed", () => {
    const msg = "AÇÕES Páginas Gerenciamento de páginas do site#titlebody1Post 1A simple page text bodyver";

    const wrapper = factory();

    expect(mockGet).toHaveBeenCalled();
    expect(wrapper.text()).toMatch(msg);
    expect(wrapper.html()).toMatchSnapshot();
  });
});
